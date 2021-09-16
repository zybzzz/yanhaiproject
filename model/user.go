/*
	包 model 存放数据库到结构体的映射
 */
package model

type User struct {
	UserId      string `gorm:"primaryKey" json:"userId"`
	Password    string `json:"password"`
	Nickname    string `json:"nickname"`
	Gender      int
	School      string `json:"school"`
	Major       string `json:"major"`
	PortraitId  int
	AttentionId string
}

//映射到数据库表名
func (User) TableName() string {
	return "user"
}