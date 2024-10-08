package model

type Group struct {
	Id   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}

func (g *Group) CreateGroup() error {
	res := DB.Create(g)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (g *Group) CheckGroupById() bool {
	res := DB.Where("id = ?", g.Id).First(g)
	if res.RowsAffected == 1 {
		return true
	}
	return false
}
