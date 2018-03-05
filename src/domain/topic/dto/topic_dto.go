package dto

type TopicDTO struct {
	Name        string `form:"Name" json:"Name" validate:"required"`
	Description string `form:"Description" json:"Description"`
}
