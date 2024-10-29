package WebsocketService

import (
	"fmt"
)

func (manager *GroupClientManager) Start() {
	for {
		fmt.Println("---------正在监听创建群聊管道通信---------")
		select {
		case gid := <-manager.Register:
			client := &GroupClient{
				GroupId:       gid,
				Broadcast:     make(chan *Broadcast),
				OnlineUserMap: make(map[string]ConnGroupClient),
			}
			manager.Clients[gid] = client
		}
	}
}
