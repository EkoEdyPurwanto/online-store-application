package req

type LoginRequest struct {
	Identifier loginIdentifier
	Password   string
}

type loginIdentifier struct {
	Email       string `json:"email" validate:"omitempty,email"`
	PhoneNumber string `json:"phoneNumber" validate:"omitempty,min=10,max=15"`
	UserName    string `json:"username" validate:"omitempty"`
}
