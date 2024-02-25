package model

import (
	"github.com/google/uuid"
	"time"
)

type Carts struct {
	Id uuid.UUID
	UserId uuid.UUID
	ProductId uuid.UUID
	Quantity int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
