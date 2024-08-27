package broneAuth

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/braciate/braciate-be/internal/entity"
	"io"
	"net/http"
	"strings"
)

func (b *broneAuth) getBetween(start, end, str string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}
	return strings.TrimSpace(str[s : s+e])
}

// makeRequest creates and executes an HTTP request, returning the response body and headers
func (b *broneAuth) makeRequest(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s request: %v", method, err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}

// parseSAMLResponse decodes and extracts student details from the SAML response
func (b *broneAuth) parseSAMLResponse(samlResponse string) (entity.User, error) {
	decodedSamlResponseByte, err := base64.StdEncoding.DecodeString(samlResponse)
	if err != nil {
		return entity.User{}, ErrUnexpected
	}

	var samlResp SAMLResponse
	err = xml.Unmarshal(decodedSamlResponseByte, &samlResp)
	if err != nil {
		return entity.User{}, ErrUnexpected
	}

	user := entity.User{}
	for _, attr := range samlResp.AttributeStatement.Attributes {
		switch attr.FriendlyName {
		case "nim":
			user.NIM = attr.AttributeValue
		case "fullName":
			user.Username = attr.AttributeValue
		case "email":
			user.Email = attr.AttributeValue
		case "fakultas":
			user.Faculty = fmt.Sprintf("Fakultas %s", attr.AttributeValue)
		case "prodi":
			user.StudyProgram = attr.AttributeValue
		}
	}

	// validation missing value
	if user.NIM == "" || user.Username == "" || user.Email == "" {
		return entity.User{}, ErrMissingRequiredFieldResponse
	}

	return user, nil
}
