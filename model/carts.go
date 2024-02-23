package model

import (
	"github.com/google/uuid"
	"time"
)

type Carts struct {
	Id uuid.UUID
	UserId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
