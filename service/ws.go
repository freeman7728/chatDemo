package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

const month = 60 * 60 * 24 * 30

type SendMSg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendId string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    string `json:"type"`
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreatId(uid, toUid string) string {
	return uid + "->" + toUid
}

func Handler(c *gin.Context) {
	uid := c.Query("uid")
	toUid := c.Query("toUid")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil) //升级websocket协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{
		ID:     CreatId(uid, toUid),
		SendId: CreatId(toUid, uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	go client.Write()
}

func (c *Client) Read() {

}
func (c *Client) Write() {

}
