package service

import (
	"chat/model"
	"chat/serializer"
)

func Insert[Model any](o Model) (resp *serializer.Response) {
	result := model.DB.Create(&o) //如果往create传any类型，则不能使用内嵌字段，因为他不知道我也没有内嵌字段
	if result.Error != nil {
		return &serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "保存失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "保存成功",
	}
}
