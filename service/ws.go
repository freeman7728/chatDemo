package service

import (
	"chat/cache"
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const month = 60 * 60 * 24 * 30

type SendMsg struct {
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
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(sendMsg)
		if err != nil {
			log.Println("数据格式不正确")
			Manager.Unregister <- c
			_ = c.Socket.Close()
			return
		}
		if sendMsg.Type == 1 {
			r1, _ := cache.RedisClient.Get(c.ID).Result()
			r2, _ := cache.RedisClient.Get(c.SendId).Result()
			if r1 >= "3" && r2 == "" {
				replyMsg := ReplyMsg{
					Code:    e.WebsocketLimit,
					Content: "达到限制",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30).Result()
				continue
			} else {
				cache.RedisClient.Incr(c.ID)
				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
			}
			log.Println(c.ID, "发送消息", sendMsg.Content)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
			}
		}
	}
}
func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Println(c.ID, "接受消息:", string(message))
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccessMessage,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
