package ws

type Trainer struct {
	Content   string `bson:"content" json:"content"`    // 内容
	StartTime int64  `bson:"startTime" json:"startime"` // 创建时间
	EndTime   int64  `bson:"endTime" json:"endTime"`    // 过期时间
	Read      uint   `bson:"read" json:"read"`          // 已读
}

type Result struct {
	StartTime int64
	From      string
	Data      Trainer `json:"data"`
}

//数据库的集合名称为senderId->receiverId
/*
	获取下线时，对方发来的消息
	{
		"type":2,
		"receiver":"1"
	}
*/
