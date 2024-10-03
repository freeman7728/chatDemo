package WebsocketService

import (
	"chat/conf"
	"chat/pkg/e"
	"chat/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func (manager *RelationClientManager) Start() {
	for {
		fmt.Println("---------正在监听关系管道通信---------")
		select {
		case conn := <-manager.Register:
			fmt.Printf("有新的连接:%v\n", conn.Uid)
			//如果有新连接接入，则插入键值对(连接id，连接)，后续方便查询用户是否在线
			RelationClientManagerIns.Clients[conn.ID] = conn
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-manager.Unregister: //断开连接
			log.Printf("连接失败:%v\n", conn.ID)
			if _, ok := RelationClientManagerIns.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(RelationClientManagerIns.Clients, conn.ID)
			}
			//监听广播通道，如果出现消息，则说明有人发送了消息
		case broadcast := <-manager.Broadcast:
			//转发到用户连接
			UserClientManagerIns.Transmit <- broadcast.Client
			//拿到接收者对象和消息体
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false
			//遍历在线用户map
			if _, flag = RelationClientManagerIns.Clients[sendId]; flag {
				select {
				//把消息体发送给接收者对象的Send通道，并且标记接收者为在线状态
				case RelationClientManagerIns.Clients[sendId].Send <- message:
					flag = true
				default:
					close(RelationClientManagerIns.Clients[sendId].Send)
					delete(RelationClientManagerIns.Clients, sendId)
				}
			}
			id := broadcast.Client.ID
			//接下来是告诉发送者本次发送的结果
			//在线
			if flag {
				log.Println("对方在线")
				replyMsg := &ReplyMsg{
					From:    broadcast.Client.Uid,
					Code:    e.WebsocketOnlineReply,
					Content: "对方在线",
				}
				msg, _ := json.Marshal(replyMsg)
				//告诉发送者，接收者在线
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//TODO 聊天记录插入数据库
				err := service.InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 1, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    e.WebsocketOfflineReply,
					Content: "对方不在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				//告诉发送者，接收者不在线
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := service.InsertMsg(conf.MongoDBName, id, string(message), 0, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			}
		}
	}
}
