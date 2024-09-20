package model

import "fmt"

type Relation struct {
	Source int64 `gorm:"uniqueIndex:idx_from_to" json:"from"` // 定义唯一索引
	Target int64 `gorm:"uniqueIndex:idx_from_to" json:"to"`   // 定义唯一索引
}

func (r *Relation) GetRelationList(id int, relationList *[]Relation) bool {
	result := DB.Where("Source = ?", id).Find(&relationList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}

func (r *Relation) CreateRelation() bool {
	res1 := DB.Create(r)
	t := r.Target
	r.Target = r.Source
	r.Source = t
	res2 := DB.Create(r)
	if res1.RowsAffected == 1 && res2.RowsAffected == 1 {
		return true
	}
	//新增对象时，对象存在，对象不存在，其他错误
	return false
}
