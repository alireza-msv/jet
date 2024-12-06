package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	httpclient "github.com/alireza-msv/jet/internal/http_client"
	"github.com/alireza-msv/jet/internal/utils"
)

type ClientOptions struct {
	AccountID    string
	ClientID     string
	ClientSecret string
	Scope        string
}

type AuthClient struct {
	ClientOptions
	accessToken     string
	tokenType       string
	restInstanceURI string
	tokenExpiresIn  *time.Time
	httpClient      *httpclient.HttpClient
}

func NewAuthClient(subdomain string, options ClientOptions) *AuthClient {
	c := AuthClient{
		ClientOptions: options,
		httpClient:    httpclient.NewHttpClient(subdomain),
	}

	if c.Scope == "" {
		c.Scope = utils.AuthDefaultScope
	}

	return &c
}

// Returns the current Access Token
// If the Access Token is expired, receives a new access token
// The new access token will be saved in the accessToken field
// and the tokenExpiresIn fileld will be updated
// The method returns empty string and error if something goes wrong
func (client *AuthClient) Token() (string, error) {
	// Returning the current token if it's not expired
	if client.tokenExpiresIn != nil && client.tokenExpiresIn.After(time.Now()) {
		return client.accessToken, nil
	}

	err := client.getToken()
	if err != nil {
		return "", err
	}

	return client.accessToken, nil
}

func (client *AuthClient) getToken() error {
	body, err := (&AuthRequest{
		GrantType:    "client_credentials",
		AccountID:    client.AccountID,
		ClientID:     client.ClientID,
		ClientSecret: client.ClientSecret,
		Scope:        client.Scope,
	}).ToJSONReader()
	if err != nil {
		return err
	}

	res, err := client.httpClient.PostJSON("v2/token", body)

	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnauthorized {
		errorResponse := AuthErrorResponse{}
		err = json.Unmarshal(resBody, &errorResponse)
		if err != nil {
			return err
		}

		return errors.New(errorResponse.ErrorDescription)
	} else if res.StatusCode != http.StatusOK {
		return errors.New(string(resBody))
	}

	authResponse := AuthResponse{}
	err = json.Unmarshal(resBody, &authResponse)
	if err != nil {
		return err
	}

	client.accessToken = authResponse.AccessToken
	client.restInstanceURI = authResponse.RESTInstanceURL
	client.tokenType = authResponse.TokenType
	expireTime := time.Now().Add(time.Duration(authResponse.ExpiresIn) * time.Second)
	client.tokenExpiresIn = &expireTime

	return nil
}

func (client *AuthClient) RESTURI() string {
	return client.restInstanceURI
}
