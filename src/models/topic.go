package models

import "time"

type Topic struct {
	ID        int        `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name        string  `gorm:"size:255" json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	News        [] News `gorm:"many2many:news_topics" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
