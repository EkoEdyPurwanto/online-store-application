package req

import "online-store-application/model"

type RegisterRequest struct {
	Role            model.Role         `json:"role"`
	Identifier      registerIdentifier `json:"identifier"`
	UserName        string             `json:"username" validate:"required,min=3,max=30"`
	Password        string             `json:"password" validate:"required"`
	PasswordConfirm string             `json:"passwordConfirm" validate:"required,eqfield=Password"`
}

type registerIdentifier struct {
	Email       string `json:"email" validate:"omitempty,email"`
	PhoneNumber string `json:"phoneNumber" validate:"omitempty,min=10,max=15"`
}
