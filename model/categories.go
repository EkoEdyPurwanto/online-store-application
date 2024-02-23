package model

import (
	"github.com/google/uuid"
	"time"
)

type Categories struct {
	Id uuid.UUID
	Type string
	CreatedAt time.Time
	UpdatedAt time.Time
}