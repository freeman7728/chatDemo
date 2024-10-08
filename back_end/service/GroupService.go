package service

import (
	"chat/model"
	"chat/serializer"
	"net/http"
	"sync"
)

var GroupServIns *GroupServ
var GroupServOnce sync.Once

func GetGroupServIns() *GroupServ {
	GroupServOnce.Do(func() {
		GroupServIns = &GroupServ{}
	})
	return GroupServIns
}

type GroupServ struct {
}

func (gs *GroupServ) CreateGroupServ(m model.Group) (resp serializer.Response) {
	err := m.CreateGroup()
	resp.Status = http.StatusOK
	if err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Error = err.Error()
		return
	}
	resp.Msg = "群聊创建成功"
	return
}
