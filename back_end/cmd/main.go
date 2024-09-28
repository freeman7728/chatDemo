package main

import (
	"chat/conf"
	"chat/router"
	"chat/service/WebsocketService"
)

func main() {
	conf.Init()
	go WebsocketService.RelationClientManagerIns.Start()
	go WebsocketService.UserClientManagerIns.Start()
	r := router.NewRouter()
	_ = r.Run("0.0.0.0" + conf.HttpPort)
}
