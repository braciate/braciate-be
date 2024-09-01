package categories

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
