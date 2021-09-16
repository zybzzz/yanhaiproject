package model

type Picture struct {
	PicId  int `gorm:"primaryKey"`
	PicURL string `gorm:"column:pic_url"`
}

func (Picture) TableName() string {
	return "picture"
}
