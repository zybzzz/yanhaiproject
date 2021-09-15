package model

type Comment struct {
	CommentId    int
	CommentusrId string
	TopicId      int
	Content      string
	picId        string
}

//映射到数据库表名
func (Comment) TableName() string {
	return "Comment"
}

