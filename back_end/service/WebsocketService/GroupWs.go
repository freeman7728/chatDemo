package WebsocketService

import (
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type GroupClient struct {
	GroupId       string
	Register      chan *ConnGroupClient
	UnRegister    chan *ConnGroupClient
	OnlineUserMap map[string]*ConnGroupClient
	Broadcast     chan *GroupBroadcast
}

type ConnGroupClient struct {
	Uid     string
	GroupId string
	Socket  *websocket.Conn
	ID      string
	Send    chan []byte
	Name    string
}

type GroupClientManager struct {
	Clients    map[string]*GroupClient
	Register   chan string
	UnRegister chan string
}

var GroupClientManagerIns = GroupClientManager{
	Clients:    make(map[string]*GroupClient),
	Register:   make(chan string),
	UnRegister: make(chan string),
}

type GroupBroadcast struct {
	Client  *ConnGroupClient
	Message []byte
	Type    string `json:"type"`
}

func GroupHandler(c *gin.Context) {
	//解析url变量
	groupId := c.Query("groupId")
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
	if _, ok := GroupClientManagerIns.Clients[groupId]; !ok {
		client := &GroupClient{
			GroupId:       groupId,
			Broadcast:     make(chan *GroupBroadcast),
			OnlineUserMap: make(map[string]*ConnGroupClient),
			Register:      make(chan *ConnGroupClient),
			UnRegister:    make(chan *ConnGroupClient),
		}
		GroupClientManagerIns.Clients[groupId] = client
		fmt.Println("---------群聊", groupId, "创建成功---------")
		go client.Start()
	}
	//创建群聊客户端对象，把连接赋予客户端对象
	client := &ConnGroupClient{
		GroupId: groupId,
		Uid:     uid,
		Socket:  conn,
		ID:      CreatId(uid, groupId),
		Send:    make(chan []byte),
	}
	GroupClientManagerIns.Clients[groupId].Register <- client //把客户端发送到在线用户通道
	//对于每一个客户端连接，都要创建conn对客户端的读写
	go client.Write()
	go client.Read()
}

// conn从客户端读
func (c *ConnGroupClient) Read() {
	defer func() {
		_ = c.Socket.Close()
		delete(GroupClientManagerIns.Clients[c.GroupId].OnlineUserMap, c.Uid)
		if len(GroupClientManagerIns.Clients[c.GroupId].OnlineUserMap) <= 0 {
			GroupClientManagerIns.UnRegister <- c.GroupId
		}
	}()
	//执行轮询
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		//解析客户端发送到服务端的消息
		err := c.Socket.ReadJSON(sendMsg)
		if err != nil {
			log.Println("用户离开群聊")
			_ = c.Socket.Close()
			return
		}
		if sendMsg.Type == 1 {
			if len(sendMsg.Content) >= 1000 {
				replyMsg := ReplyMsg{
					Code:    e.WebsocketLimit,
					Content: "消息长度过长",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			log.Println(c.ID, "群聊发送消息", sendMsg.Content)
			//TODO 创建broadcast对象，送过去
			broadcast := &GroupBroadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
				Type:    "1",
			}
			GroupClientManagerIns.Clients[c.GroupId].Broadcast <- broadcast
		} else if sendMsg.Type == 2 { //拉取历史消息
			//TODO

		}
	}
}

// conn往客户端写
func (c *ConnGroupClient) Write() {
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
