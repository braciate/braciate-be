package nominations

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateNominationRequest struct {
	Name       string `json:"name"`
	CategoryID string `json:"categories_id"`
}

type CreateNominationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"categories_id"`
}
