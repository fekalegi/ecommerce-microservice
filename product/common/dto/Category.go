package dto

import "github.com/google/uuid"

type RequestCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type RequestUpdateCategory struct {
	Name string `json:"name"`
}
