package WebsocketService

//
//import (
//	"chat/conf"
//	"chat/pkg/e"
//	"chat/service"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//)
//
//type GroupClient struct {
//	Uid     string
//	GroupId string
//	Socket  *websocket.Conn
//	ID      string
//	Send    chan []byte
//	Name    string
//}
//
//func GroupHandler(c *gin.Context) {
//	//解析url变量
//	uid := c.Query("groupId")
//	groupId := c.Query("uid")
//	conn, err := (&websocket.Upgrader{ //新建websocket对象
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}).Upgrade(c.Writer, c.Request, nil) //调用Upgrade方法
//	if err != nil {
//		http.NotFound(c.Writer, c.Request)
//		return
//	}
//	//创建群聊客户端对象，把连接赋予客户端对象
//	client := &GroupClient{
//		Uid:     uid,
//		GroupId: groupId,
//		Socket:  conn,
//		Send:    make(chan []byte),
//		Name: uid + "=" + groupId,
//	}
//	//RelationClientManagerIns.Register <- client //把客户端发送到在线用户通道
//	//对于每一个客户端连接，都要创建conn对客户端的读写
//	go client.Write()
//	go client.Read()
//}
//
//// conn从客户端读
//func (c *GroupClient) Read() {
//	defer func() {
//		RelationClientManagerIns.Unregister <- c
//		_ = c.Socket.Close()
//	}()
//	//执行轮询
//	for {
//		c.Socket.PongHandler()
//		sendMsg := new(SendMsg)
//		//解析客户端发送到服务端的消息
//		err := c.Socket.ReadJSON(sendMsg)
//		if err != nil {
//			log.Println("数据格式不正确")
//			_ = c.Socket.Close()
//			return
//		}
//		if sendMsg.Type == 1 {
//			if len(sendMsg.Content) >= 1000 {
//				replyMsg := ReplyMsg{
//					Code:    e.WebsocketLimit,
//					Content: "消息长度过长",
//				}
//				msg, _ := json.Marshal(replyMsg)
//				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
//				continue
//			}
//			//r1, _ := cache.RedisClient.Get(c.ID).Result()
//			//r2, _ := cache.RedisClient.Get(c.SendID).Result()
//			//if r1 >= "3" && r2 == "" {
//			//	replyMsg := ReplyMsg{
//			//		Code:    e.WebsocketLimit,
//			//		Content: "达到限制",
//			//	}
//			//	msg, _ := json.Marshal(replyMsg)
//			//	_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
//			//	_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30).Result()
//			//	continue
//			//} else {
//			//	cache.RedisClient.Incr(c.ID)
//			//	_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
//			//}
//			log.Println(c.ID, "发送消息", sendMsg.Content)
//			//取出ReceiverId字段，发送到broadcast通道
//			RelationClientManagerIns.Broadcast <- &Broadcast{
//				Client:  c,
//				Message: []byte(sendMsg.Content),
//			}
//		} else if sendMsg.Type == 2 { //拉取历史消息
//			results, _ := service.FindMany(conf.MongoDBName, c.SendID, c.ID)
////			if len(results) == 0 {
////				replyMsg := ReplyMsg{
////					Code:    e.WebsocketEnd,
////					Content: "到底了",
////				}
////				msg, _ := json.Marshal(replyMsg)
////				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////				continue
////			}
////			for _, result := range results {
////				//封装消息，使历史消息与正常接受发送的消息格式一致
////				msg, _ := json.Marshal(ReplyMsg{
////					Code:    50001,
////					From:    result.From,
////					Content: result.Data.Content,
////				})
////				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////			}
////		}
//	}
//}
//
//// conn往客户端写
//func (c *GroupClient) Write() {
//	defer func() {
//		_ = c.Socket.Close()
//	}()
//	for {
//		select {
//		//作为接收者，监听send通道，如果有则取出
//		case message, ok := <-c.Send:
//			if !ok {
//				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
//				return
//			}
//			log.Println(c.ID, "接受消息:", string(message))
//			replyMsg := ReplyMsg{
//				Code: e.WebsocketSuccessMessage,
//				//fmt包格式化字符串
//				Content: fmt.Sprintf("%s", string(message)),
//			}
//			msg, _ := json.Marshal(replyMsg)
//			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
//		}
//	}
//}
