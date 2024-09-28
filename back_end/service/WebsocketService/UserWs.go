package WebsocketService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type UserClient struct {
	Uid    string
	Socket *websocket.Conn
	Send   chan []byte
}

type UserClientManager struct {
	Clients        map[string]*UserClient
	UserRegister   chan *UserClient
	UserUnregister chan *UserClient
	Transmit       chan *Client
}

var UserClientManagerIns = UserClientManager{
	Clients:        make(map[string]*UserClient),
	UserRegister:   make(chan *UserClient),
	UserUnregister: make(chan *UserClient),
	Transmit:       make(chan *Client),
}

func UserWsHandler(c *gin.Context) {
	//TODO id查不到
	id, _ := c.Get("id")
	conn, err := (&websocket.Upgrader{ //新建websocket对象
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil) //调用Upgrade方法
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &UserClient{
		Uid:    fmt.Sprintf("%v", id),
		Socket: conn,
		Send:   make(chan []byte),
	}
	UserClientManagerIns.UserRegister <- client
	go client.Write()
}

func (c *UserClient) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {

			}
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}

		}
	}
}
