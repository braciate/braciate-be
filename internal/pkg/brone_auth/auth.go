package broneAuth

import (
	"fmt"
	"github.com/braciate/braciate-be/internal/entity"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type broneAuth struct {
	client *http.Client
}

type BroneAuth interface {
	Authenticate(identifier, password string) (entity.User, error)
}

func New() BroneAuth {
	return &broneAuth{
		client: &http.Client{},
	}
}

func (b *broneAuth) Authenticate(identifier, password string) (entity.User, error) {
	initialHeaders := map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language":           "en-US,en;q=0.9,id;q=0.8",
		"cache-control":             "no-cache",
		"pragma":                    "no-cache",
		"referer":                   "https://brone.ub.ac.id/",
		"sec-ch-ua":                 `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"macOS"`,
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	}

	resp, err := b.makeRequest("GET", "https://brone.ub.ac.id/my/", initialHeaders, nil)
	if err != nil {
		return entity.User{}, ErrMakingRequestToBrone
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.User{}, ErrUnexpected
	}

	r := string(respBody)
	cookies := fmt.Sprintf("%s", resp.Header["Set-Cookie"])
	authSessionID := b.getBetween("AUTH_SESSION_ID=", ";", cookies)
	authSessionIDLegacy := b.getBetween("AUTH_SESSION_ID_LEGACY=", ";", cookies)
	kcRestart := b.getBetween("KC_RESTART=", ";", cookies)

	fullURL := b.getBetween(`action="`, `" `, r)
	sessionCode := b.getBetween("session_code=", "&amp", fullURL)
	execution := b.getBetween("execution=", "&amp", fullURL)
	tabID := strings.Split(fullURL, "tab_id=")[1]

	postData := url.Values{
		"username":     {identifier},
		"password":     {password},
		"credentialId": {""},
	}
	postHeaders := map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language":           "en-US,en;q=0.9,id;q=0.8",
		"cache-control":             "no-cache",
		"content-type":              "application/x-www-form-urlencoded",
		"pragma":                    "no-cache",
		"sec-ch-ua":                 `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"macOS"`,
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		"cookie":                    fmt.Sprintf("AUTH_SESSION_ID=%s; AUTH_SESSION_ID_LEGACY=%s; KC_RESTART=%s", authSessionID, authSessionIDLegacy, kcRestart),
	}

	resp, err = b.makeRequest("POST", fmt.Sprintf("https://iam.ub.ac.id/auth/realms/ub/login-actions/authenticate?session_code=%s&execution=%s&client_id=brone.ub.ac.id&tab_id=%s", sessionCode, execution, tabID), postHeaders, strings.NewReader(postData.Encode()))
	if err != nil {
		return entity.User{}, ErrMakingRequestToIam
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return entity.User{}, ErrUnexpected
	}

	r = string(respBody)
	if !strings.Contains(r, "SAMLResponse") {
		if strings.Contains(r, "Invalid username or password.") {
			return entity.User{}, ErrInvalidEmailNimOrPassword
		}
		return entity.User{}, ErrResponseNotContainsSAML
	}

	samlResponse := b.getBetween(`name="SAMLResponse" value="`, `"/>`, r)
	return b.parseSAMLResponse(samlResponse)
}
