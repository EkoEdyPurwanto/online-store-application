package model

import (
	"github.com/google/uuid"
	"time"
)

type Products struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	ProductName string
	Price       int64
	Stock       int64
	CategoryId  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
