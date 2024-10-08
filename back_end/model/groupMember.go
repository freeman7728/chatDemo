package model

type GroupMember struct {
	UserId  uint `json:"user_id"`
	GroupId uint `json:"group_id"`
}

func (g *GroupMember) AddGroupMember() error {
	res := DB.Create(g)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
