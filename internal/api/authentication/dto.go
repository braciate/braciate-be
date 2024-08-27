package authentication

type SigninRequest struct {
	NimEmail string `json:"nim_email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SigninResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   int64  `json:"expired_at"`
}
