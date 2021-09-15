package model

type Group struct {
	GroupId     int
	GroupSchool string
	GroupPicId  int
}

//映射到数据库表名
func (Group) TableName() string {
	return "Group"
}
