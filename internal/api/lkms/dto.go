package lkms

type LkmsRequest struct {
	Name       string
	CategoryID string
	Type       int
}

type LkmsResponse struct {
	ID         string
	Name       string
	CategoryID string
	LogoLink   string
	Type       int
}
