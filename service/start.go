package service

import (
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func (manager *ClientManager) Start() {
	for {
		fmt.Println("---------正在监听管道通信---------")
		select {
		case conn := <-manager.Register:
			fmt.Printf("有新的连接:%v\n", conn.ID)
			Manager.Clients[conn.ID] = conn
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-manager.Unregister: //断开连接
			log.Printf("连接失败:%v\n", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendId
			flag := false
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, sendId)
				}
			}
			//id := broadcast.Client.ID
			//在线
			if flag {
				log.Println("对方在线")
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketOnlineReply,
					Content: "对方在线",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//TODO 聊天记录插入数据库
				//err = InsertMsg(conf.MongoDBName, id, string(message), 1, int64(3*month))
				//if err != nil {
				//	fmt.Println("InsertOneMsg Err", err)
				//}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    e.WebsocketOfflineReply,
					Content: "对方不在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//err = InsertMsg(conf.MongoDBName, id, string(message), 0, int64(3*month))
				//if err != nil {
				//	fmt.Println("InsertOneMsg Err", err)
				//}
			}
			//不在线
		}

		//TODO 消息的广播/传递

	}
}
