package auth

import (
	"bytes"
	"encoding/json"
	"io"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
	AccountID    string `json:"account_id"`
}

func (ar *AuthRequest) ToJSON() ([]byte, error) {
	return json.Marshal(ar)
}

// Converts the struct into json and creates a io.Reader from it
func (ar *AuthRequest) ToJSONReader() (io.Reader, error) {
	b, err := ar.ToJSON()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

type AuthResponse struct {
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	TokenType       string `json:"token_type"`
	RESTInstanceURL string `json:"rest_instance_url"`
	SOATInstanceURL string `json:"soap_instance_url"`
	Scope           string `json:"scope"`
}

type AuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}
