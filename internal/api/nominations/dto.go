package nominations

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NominationRequest struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

type NominationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}
