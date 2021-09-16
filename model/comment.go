package model

type Comment struct {
	CommentId    int `gorm:"primaryKey"`
	CommentusrId string
	TopicId      int
	Content      string
	PicId        string
}

//映射到数据库表名
func (Comment) TableName() string {
	return "comment"
}

