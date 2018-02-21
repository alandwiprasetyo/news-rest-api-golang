package dto

type NewsDTO struct {
	Headline    string `form:"Headline" json:"headline" validate:"required"`
	Title       string `form:"Title" json:"title" validate:"required"`
	Description string `form:"Description" json:"Description" validate:"required"`
	Tags        string `form:"Tags" json:"tags"`
	Status      string `form:"Status" json:"status"`
	Type        string `sql:"not null;type:ENUM('publish', 'draft', 'deleted');default:'draft'"`
}
