package WebsocketService

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (manager *GroupClientManager) Start() {
	for {
		fmt.Println("---------正在监听创建群聊管道通信---------")
		select {
		case gid := <-manager.UnRegister:
			fmt.Println("---------群聊", gid, "销毁---------")
			delete(manager.Clients, gid)
		}
	}
}

func (manager *GroupClient) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.OnlineUserMap[conn.Uid] = conn
			log.Info("用户 ", conn.Uid, " 加入 ", conn.GroupId, " 群聊", "目前", len(manager.OnlineUserMap), "人")
		case broadcast := <-manager.Broadcast:
			for _, conn := range manager.OnlineUserMap {
				conn.Send <- broadcast
			}
		}
	}
}
