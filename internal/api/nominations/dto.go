package nominations

type CreateNominationRequest struct {
	Name       string `json:"name"`
	CategoryID string `json:"categories_id"`
}

type CreateNominationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"categories_id"`
}
