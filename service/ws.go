package service

import (
	"chat/cache"
	"chat/conf"
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

const month = 60 * 60 * 24 * 30

type SendMsg struct {
	Type       int    `json:"type"`
	Content    string `json:"content"`
	ReceiverId string `json:"receiver"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	Uid    string
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client     *Client
	Message    []byte
	Type       string `json:"type"`
	ReceiverId string `json:"receiver"`
}

type ClientManager struct {
	Clients   map[string]*Client
	Broadcast chan *Broadcast
	//Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager 包级变量，执行main函数之前就会进行初始化
var Manager = ClientManager{
	Clients:   make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast: make(chan *Broadcast),
	Register:  make(chan *Client),
	//Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreatId(uid, toUid string) string {
	return uid + "->" + toUid
}

func Handler(c *gin.Context) {
	//解析url变量
	uid := c.Query("uid")
	conn, err := (&websocket.Upgrader{ //新建websocket对象
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil) //调用Upgrade方法
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//创建客户端对象，把连接赋予客户端对象
	client := &Client{
		Uid:    uid,
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client //把客户端发送到在线用户通道
	//对于每一个客户端连接，都要创建conn对客户端的读写
	go client.Read()
	go client.Write()
}

/*
	把客户端发送来的消息{
		发送者，
		接收者，
		消息体
	}转发到broadcast通道中，再由broadcast通道去寻找接收者，观察其是否在线
*/

// conn从客户端读
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	//执行轮询
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		//解析客户端发送到服务端的消息
		err := c.Socket.ReadJSON(sendMsg)
		if err != nil {
			log.Println("数据格式不正确")
			Manager.Unregister <- c
			_ = c.Socket.Close()
			return
		}
		if sendMsg.Type == 1 {
			r1, _ := cache.RedisClient.Get(c.ID).Result()
			r2, _ := cache.RedisClient.Get(c.ID).Result()
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
			//取出ReceiverId字段，发送到broadcast通道
			Manager.Broadcast <- &Broadcast{
				Client:     c,
				Message:    []byte(sendMsg.Content),
				ReceiverId: sendMsg.ReceiverId,
			}
		} else if sendMsg.Type == 2 { //拉取历史消息
			timeT, err := strconv.Atoi(sendMsg.Content) // 传送来时间
			if err != nil {
				timeT = 999999999
			}
			c.ID = CreatId(c.Uid, sendMsg.ReceiverId)
			c.SendID = CreatId(sendMsg.ReceiverId, c.Uid)
			results, _ := FindMany(conf.MongoDBName, c.SendID, c.ID, int64(timeT), 10)
			if len(results) > 10 {
				results = results[:10]
			} else if len(results) == 0 {
				replyMsg := ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "到底了",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			for _, result := range results {
				msg, _ := json.Marshal(result)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

// conn往客户端写
func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		//作为接收者，监听send通道，如果有则取出
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Println(c.ID, "接受消息:", string(message))
			replyMsg := ReplyMsg{
				Code: e.WebsocketSuccessMessage,
				//fmt包格式化字符串
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
