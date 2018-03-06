package tables

import "github.com/alandwiprasetyo/rest-api/src/models/base"

const (
	PUBLISH string = "publish"
	DRAFT   string = "draft"
	DELETED string = "deleted"
)

type News struct {
	base.Model
	Headline    string  `json:"headline"`
	Title       string  `gorm:"size:255" json:"title"`
	Status      string  `gorm:"size:255;default:'draft'" sql:"not null;type:ENUM('draft', 'publish', 'deleted');default:'draft'"  json:"status"`
	Description string  `json:"description"`
	Tags        string  `json:"tags"`
	Topic       []Topic `gorm:"many2many:news_topics" json:"topics"`
}