package assets

type AssetsRequest struct {
	UserID       string `json:"user_id" validate:"required"`
	LkmID        string `json:"lkm_id" validate:"required"`
	NominationID string `json:"nomination_id" validate:"required"`
	Url          string `json:"url" validate:"required"`
}

type AssetsResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	LkmID        string `json:"lkm_id"`
	NominationID string `json:"nomination_id"`
	Url          string `json:"url"`
}
