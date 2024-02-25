package req

import "github.com/google/uuid"

type CreateProductsRequest struct {
	UserId      uuid.UUID `json:"user_id"`
	ProductName string    `json:"product_name" validate:"required,min=3,max=20"`
	Price       int64     `json:"price" validate:"required,gt=0"`
	Stock       int64     `json:"stock" validate:"required,gt=0"`
	CategoryId  uuid.UUID `json:"category_id" validate:"required"`
}
