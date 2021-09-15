package model

type Picture struct {
	PicId  int
	PicURL string
}

func (Picture) TableName() string {
	return "Picture"
}
