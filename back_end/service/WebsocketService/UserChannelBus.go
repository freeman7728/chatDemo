package WebsocketService

import (
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func (manager *UserClientManager) Start() {
	for {
		fmt.Println("---------正在监听用户管道通信---------")
		select {
		case client := <-manager.UserRegister:
			log.Info("----------有新的用户连接----------")
			UserClientManagerIns.Clients[client.Uid] = client
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
		case relationClient := <-manager.Transmit:
			//告诉对应用户有新的会话
			if _, ok := UserClientManagerIns.Clients[relationClient.ToUid]; !ok {
				log.Info("对方不在线")
			} else {
				msg := UserClientMsg{Uid: relationClient.Uid}
				UserClientManagerIns.Clients[relationClient.ToUid].Send <- msg
				log.Info("对方在线")
			}
		case client := <-manager.UserUnregister:
			replyMsg := &ReplyMsg{
				Code:    e.WebsocketEnd,
				Content: "连接已断开",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
			close(client.Send)
			delete(UserClientManagerIns.Clients, client.Uid)
			log.Info("用户", client.Uid, "登出")
		}
	}
}
