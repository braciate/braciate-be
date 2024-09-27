package userVotes

type UserVotesRequest struct {
	UserID       string `json:"user_id" validate:"required"`
	LkmID        string `json:"lkm_id" validate:"required"`
	NominationID string `json:"nomination_id" validate:"required"`
}

type UserVotesResponse struct {
	ID           string `json:"id" validate:"required"`
	UserID       string `json:"user_id" validate:"required"`
	LkmID        string `json:"lkm_id" validate:"required"`
	NominationID string `json:"nomination_id" validate:"required"`
}
