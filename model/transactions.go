package model

import (
	"github.com/google/uuid"
	"time"
)

type Transactions struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	ProductId  uuid.UUID
	Quantity   int64
	TotalPrice int64
	Date       time.Time
}
