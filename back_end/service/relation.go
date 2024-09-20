package service

import (
	"chat/model"
	"chat/serializer"
	"net/http"
	"sync"
)

type RelationServ struct {
}
type RelationList struct {
	List []model.Relation `json:"list"`
}

var RelationServIns *RelationServ
var RelationServOnce sync.Once

func GetRelationServIns() *RelationServ {
	RelationServOnce.Do(func() {
		RelationServIns = &RelationServ{}
	})
	return RelationServIns
}

// CreateRelation 从JWT中得到发送者的id
// CreateRelation 检查双方id，发送者的id与登录账号需要匹配，接收者的id需要存在
func (r *RelationServ) CreateRelation(relation model.Relation) (resp serializer.Response) {
	resp.Status = http.StatusOK
	var user model.User
	if !user.CheckUid(uint(relation.Target)) {
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
func (r *RelationServ) GetRelation(id int64) (resp serializer.Response) {
	resp.Status = http.StatusOK
	var relations model.Relation
	relationList := make([]model.Relation, 0)
	if !relations.GetRelationList(int(id), &relationList) {
		resp.Status = http.StatusInternalServerError
		resp.Msg = "查询出错"
		return
	}
	if len(relationList) == 0 {
		resp.Msg = "好友列表为空"
		return
	}
	var list RelationList
	list.List = relationList
	resp.Data = list
	resp.Msg = "success"
	return
}
