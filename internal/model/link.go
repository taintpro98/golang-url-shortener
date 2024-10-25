package model

type LinkModel struct {
	ID          int64
	Short       string
	OriginalURL string
}

func (LinkModel) TableName() string {
	return "links"
}
