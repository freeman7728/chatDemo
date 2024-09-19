package service

import (
	"chat/model"
	"chat/serializer"
	"net/http"
	"sync"
)

type Relation struct {
}

var RelationServIns *Relation
var RelationServOnce sync.Once

func GetRelationServIns() *Relation {
	RelationServOnce.Do(func() {
		RelationServIns = &Relation{}
	})
	return RelationServIns
}

// CreateRelation 从JWT中得到发送者的id
// CreateRelation 检查双方id，发送者的id与登录账号需要匹配，接收者的id需要存在
func (r *Relation) CreateRelation(relation model.Relation) (resp serializer.Response) {
	resp.Status = http.StatusOK
	var user model.User
	if !user.CheckUid(uint(relation.To)) {
		resp.Status = http.StatusInternalServerError
		resp.Msg = "对方不存在"
		return
	}
	if relation.CreateRelation() {
		resp.Msg = "关系建立成功"
	} else {
		resp.Msg = "不能重复添加好友"
		resp.Status = http.StatusInternalServerError
	}
	return
}
