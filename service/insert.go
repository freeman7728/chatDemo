package service

import (
	"chat/model"
	"chat/serializer"
)

type CommonCRUDService[T any] struct {
}

func NewCommonCRUDService[T any]() CommonCRUDService[T] {
	return CommonCRUDService[T]{}
}

func (c *CommonCRUDService[T]) Insert(o T) (resp *serializer.Response) {
	result := model.DB.Create(&o)
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
