package service

import (
	"chat/model"
	"chat/pkg/e"
	"chat/pkg/utils"
	"chat/serializer"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

type UserLoginVo struct {
	Token string `json:"token"`
}

// UserRegisterService 用户注册
func (u *UserService) UserRegisterService() (resp serializer.Response) {
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

// Login 用户登录
func (u *UserService) Login() (resp serializer.Response) {
	user := &model.User{
		UserName: u.UserName,
	}
	if user.CheckPassword(u.PassWord) == true {
		user.GetUserIdByUserName()
		token, err := utils.GenerateToken(int64(user.ID), user.UserName)
		if err != nil {
			return serializer.Response{
				Status: e.ERROR, Msg: "token生成失败", Data: nil,
			}
		}
		return serializer.Response{
			Status: e.SUCCESS,
			Msg:    e.GetMsg(e.SUCCESS),
			Data:   UserLoginVo{Token: token},
		}
	}
	return serializer.Response{
		Status: e.ERROR, Msg: "账号或者密码错误",
	}
}
