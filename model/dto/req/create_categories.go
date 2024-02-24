package req

type CreateCategoriesRequest struct {
	Type string `json:"type" validate:"required,min=3,max=10"`
}
