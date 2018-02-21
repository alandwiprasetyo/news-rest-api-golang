package models

import "time"

type Model struct {
	ID        int        `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at" sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt time.Time  `json:"updated_at" sql:"DEFAULT:'current_timestamp'"`
	DeletedAt *time.Time `json:"deleted_at"`
}