package service

import (
	"chat/model"
	"chat/pkg/e"
	"chat/serializer"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

func (u UserService) UserRegisterService() (resp serializer.Response) {
	//TODO 参数校验，然后传递给持久化
	/*
		TODO
		确认username是否存在，若存在则返回
		不存在则加密密码且创建账号
	*/
	var user model.User
	var count int
	code := e.SUCCESS
	model.DB.Model(&model.User{}).Where("user_name=?", u.UserName).Count(&count)
	if count == 1 {
		code = e.ERROR
		log.Info("用户", u.UserName, "已存在")
		return serializer.Response{
			Status: code, Data: e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: u.UserName,
		Status:   model.Active,
	}
	//加密密码
	if err := user.SetPassword(u.PassWord); err != nil {
		code = e.ERROR
		log.Error(err)
		return serializer.Response{Status: code, Data: "密码加密失败"}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		code = e.ERROR
		log.Error(err)
		return serializer.Response{Status: code, Data: e.GetMsg(code)}
	}
	return serializer.Response{
		Status: code, Msg: e.GetMsg(code),
	}
}
