package models

type Command struct {
	Operation string `json:"operation" validate:"required"`
}
