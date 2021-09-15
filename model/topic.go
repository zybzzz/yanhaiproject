package model

import "time"

//FIXME 存在时间格式转化问题
type Topic struct {
	TopicId    int
	GroupId    int
	Creator    string
	Title      string
	Content    string
	CreateTime time.Time
	ThumpUp    int
	Tag        string
	School     string
	Major      string
	PicId      string
}

//映射到数据库表名
func (Topic) TableName() string {
	return "Topic"
}
