package model

type Group struct {
	GroupId     int `gorm:"primaryKey"`
	GroupSchool string
	GroupPicId  int
}

//映射到数据库表名
func (Group) TableName() string {
	return "group"
}
