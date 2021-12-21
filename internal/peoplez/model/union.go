package model

import "github.com/google/uuid"

type Union struct {
	ID          uuid.UUID
	Name        string
	Aim         string
	Contact     string
	Manager     uuid.UUID
	Description string
}
