package nominations

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateNominationRequest struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

type NominationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}
