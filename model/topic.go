package model

import "time"

//FIXME 存在时间格式转化问题
type Topic struct {
	TopicId  int `gorm:"primaryKey"`
	GroupId  int
	Creator  string
	Title    string
	Content  string
	CreateAt time.Time `gorm:"column:create_time"`
	ThumpUp  int
	Tag      string
	School   string
	Major    string
	PicId    string
}

//映射到数据库表名
func (Topic) TableName() string {
	return "topic"
}
