package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
	AccountID    string `json:"account_id"`
}

type AuthResponse struct {
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	TokenType       string `json:"token_type"`
	RESTInstanceURL string `json:"rest_instance_url"`
	SOATInstanceURL string `json:"soap_instance_url"`
	Scope           string `json:"scope"`
}

func DoAuth(subDomain string, authRequest *AuthRequest) (*AuthResponse, error) {
	requestURL := fmt.Sprintf("https://%s.auth.marketingcloudapis.com", subDomain)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

}
