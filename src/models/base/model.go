package base

import "time"

type Model struct {
	ID        int        `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}