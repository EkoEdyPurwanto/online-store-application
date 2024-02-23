package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	UsersStatus string
	Role        string
)

const (
	Active   UsersStatus = "active"
	Inactive UsersStatus = "inactive"
	Blocked  UsersStatus = "blocked"
	Other    UsersStatus = "other"
	Admin    Role        = "admin"
	Merchant Role        = "merchant"
	Customer Role        = "customer"
)

type Users struct {
	Id          uuid.UUID
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	UserStatus  UsersStatus
	Role        Role
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
