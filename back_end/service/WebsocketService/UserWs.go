package WebsocketService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type UserClientMsg struct {
	Uid string `json:"uid"`
}

type UserClient struct {
	Uid    string
	Socket *websocket.Conn
	Send   chan UserClientMsg
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
		Send:   make(chan UserClientMsg),
	}
	UserClientManagerIns.UserRegister <- client
	go client.Write()
	go client.Read()
}

func (c *UserClient) Read() {
	defer func() {
		_ = c.Socket.Close()
		UserClientManagerIns.UserUnregister <- c
		fmt.Println("logout")
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		//解析客户端发送到服务端的消息
		err := c.Socket.ReadJSON(sendMsg)
		if err != nil {
			return
		}
	}
}

func (c *UserClient) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case m, ok := <-c.Send:
			if !ok {

			}
			msg, err := json.Marshal(m)
			if err != nil {
				return
			}
			err = c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}

		}
	}
}
