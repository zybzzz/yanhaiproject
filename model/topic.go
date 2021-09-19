package model

import "time"

//FIXME 存在时间格式转化问题
type Topic struct {
	TopicId  int `gorm:"primaryKey;AUTO_INCREMENT"`
	GroupId  int `json:"groupId"`
	Creator  string `json:"creator"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt time.Time `gorm:"column:create_time"`
	ThumpUp  int
	Tag      string `json:"tag"`
	School   string
	Major    string
	PicId    string `json:"picId"`
}

//映射到数据库表名
func (Topic) TableName() string {
	return "topic"
}
