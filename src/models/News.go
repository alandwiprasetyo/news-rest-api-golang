package models

const (
	PUBLISH string = "publish"
	DRAFT   string = "draft"
	DELETED string = "deleted"
)

type News struct {
	Model
	Headline    string  `json:"headline"`
	Title       string  `gorm:"size:255" json:"title"`
	Status      string  `gorm:"size:255;default:'draft'"`
	Description string  `json:"description"`
	Tags        string  `json:"tags"`
	Topic       []Topic `gorm:"many2many:news_topics" json:"topics"`
}
