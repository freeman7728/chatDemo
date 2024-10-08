package service

import (
	"chat/model"
	"chat/serializer"
	"net/http"
	"sync"
)

type GroupMemberServ struct{}

var GroupMemberServIns *GroupMemberServ
var GroupMemberServOnce sync.Once

func GetGroupMemberServIns() *GroupMemberServ {
	GroupMemberServOnce.Do(func() {
		GroupMemberServIns = &GroupMemberServ{}
	})
	return GroupMemberServIns
}

func (g *GroupMemberServ) AddGroupMember(m model.GroupMember) (resp serializer.Response) {
	//检查用户和群是否存在
	var user model.User
	var group model.Group
	group.Id = m.GroupId
	if !user.CheckUid(m.UserId) || !group.CheckGroupById() {
		resp.Status = http.StatusInternalServerError
		resp.Error = "用户或群聊不存在"
		return
	}
	err := m.AddGroupMember()
	resp.Status = http.StatusOK
	if err != nil {
		resp.Status = http.StatusInternalServerError
		//resp.Error = err.Error()
		resp.Msg = "不可重复加入群聊"
		return
	}
	resp.Msg = "群聊加入成功"
	return
}
