/*
	包 model 存放数据库到结构体的映射
 */
package model

type User struct {
	UserId      string
	Password    string
	Nickname    string
	Gender      int
	School      string
	Major       string
	PortraitId  int
	AttentionId string
}

//映射到数据库表名
func (User) TableName() string {
	return "User"
}