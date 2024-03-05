package domain

import "github.com/google/uuid"

type Pool struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}
