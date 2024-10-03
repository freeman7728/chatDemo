package service

import (
	"chat/conf"
	"chat/model"
	"chat/serializer"
	"net/http"
	"slices"
	"strconv"
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
	if relation.Target == relation.Source {
		resp.Status = http.StatusInternalServerError
		resp.Msg = "不能添加自己"
		return
	}
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

// DelRelation 删除好友
func (r *RelationServ) DelRelation(relation model.Relation) (resp serializer.Response) {
	resp.Status = http.StatusOK
	if relation.Target == relation.Source {
		resp.Status = http.StatusInternalServerError
		resp.Msg = "不能删除自己"
		return
	}
	if relation.DelRelation() {
		resp.Msg = "关系删除成功"
	} else {
		resp.Msg = "检查好友是否存在"
		resp.Status = http.StatusInternalServerError
	}
	return
}

// GetRelation  获取好友列表
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
	sortedlist := make([]LatestUser, 0)
	for _, relation := range relationList {
		sortedlist = append(sortedlist, FindLatest(conf.MongoDBName, strconv.FormatInt(relation.Source, 10), strconv.FormatInt(relation.Target, 10)))
	}
	slices.SortFunc(sortedlist, func(a, b LatestUser) int {
		return int(b.StartTime - a.StartTime)
	})
	resList := make([]model.Relation, 0)
	for _, s := range sortedlist {
		parseInt, _ := strconv.ParseInt(s.Uid, 10, 64)
		resList = append(resList, model.Relation{Source: id, Target: parseInt})
	}
	var list RelationList
	list.List = resList
	resp.Data = list
	resp.Msg = "success"
	return
}
