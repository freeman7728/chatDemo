package WebsocketService

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (manager *UserClientManager) Start() {
	for {
		select {
		case client := <-UserClientManagerIns.UserRegister:
			log.Info("----------有新的用户连接----------")
			UserClientManagerIns.Clients[client.Uid] = client
			fmt.Println(client.Uid)
		case relationClient := <-UserClientManagerIns.Transmit:
			//告诉对应用户有新的会话
			if _, ok := UserClientManagerIns.Clients[relationClient.ToUid]; !ok {
				log.Info("对方不在线")
				fmt.Println(relationClient.ToUid)
			} else {
				log.Info("对方在线")
			}
		}
	}
}
