package req

import "github.com/google/uuid"

type CreateCartsRequest struct {
	UserId uuid.UUID `json:"user_id"`
	ProductId uuid.UUID `json:"product_id"`
	Quantity int64 `json:"quantity" validate:"required,gt=0"`
}
