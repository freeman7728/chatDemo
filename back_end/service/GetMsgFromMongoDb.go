package service

import (
	"chat/conf"
	"chat/model/ws"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}
type LatestUser struct {
	Uid       string `json:"uid"`
	StartTime int64  `json:"start_time"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDbClient.Database(database).Collection(id)
	comment := ws.TrainerWriter{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}

func FindMany(database string, sendId string, id string) (results []ws.Result, err error) {
	var resultsMe []ws.TrainerReader
	var resultsYou []ws.TrainerReader
	sendIdCollection := conf.MongoDbClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDbClient.Database(database).Collection(id)
	// 如果不知道该使用什么context，可以通过context.TODO() 产生context
	//sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(),
	//	options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	//idTimeCursor, err := idCollection.Find(context.TODO(),
	//	options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	//TODO 消息记录读取后改变消息读取状态，以及按照发送顺序进行排序
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(),
		bson.D{}, // 查询条件，可以根据需要修改,
		options.Find().SetSort(bson.D{{"startTime", 1}}),
	)
	idTimeCursor, err := idCollection.Find(context.TODO(),
		bson.D{}, // 查询条件，可以根据需要修改,
		options.Find().SetSort(bson.D{{"startTime", 1}}),
	)
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}

func FindLatest(database, id, toUid string) (l LatestUser) {
	var resultsMe []ws.TrainerReader
	var resultsYou []ws.TrainerReader
	sendIdCollection := conf.MongoDbClient.Database(database).Collection(toUid + "->" + id)
	idCollection := conf.MongoDbClient.Database(database).Collection(id + "->" + toUid)
	sendIdTimeCursor, _ := sendIdCollection.Find(context.TODO(),
		bson.D{}, // 查询条件，可以根据需要修改,
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(1),
	)
	idTimeCursor, _ := idCollection.Find(context.TODO(),
		bson.D{}, // 查询条件，可以根据需要修改,
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(1),
	)
	_ = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	_ = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	var yt int64
	var mt int64
	if len(resultsYou) > 0 {
		yt = resultsYou[0].StartTime
	}
	if len(resultsMe) > 0 {
		mt = resultsMe[0].StartTime
	}
	if yt > mt {
		l.Uid = toUid
		l.StartTime = yt
	} else {
		l.Uid = toUid
		l.StartTime = mt
	}
	return
}

func AppendAndSort(resultsMe, resultsYou []ws.TrainerReader) (results []ws.Result, err error) {
	for _, r := range resultsMe {
		result := ws.Result{
			StartTime: r.StartTime,
			From:      "me",
			Data:      r,
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		result := ws.Result{
			StartTime: r.StartTime,
			From:      "you",
			Data:      r,
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}
