package lkms

type LkmsRequest struct {
	Name       string `json:"id"`
	CategoryID string `json:"category_id"`
	Type       int    `json:"type"`
}

type LkmsResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	LogoFile   string `json:"logo_file"`
	Type       int    `json:"type"`
}
