package salesforce

import (
	"github.com/alireza-msv/jet/internal/auth"
)

type SalesForceClient struct {
	authClient *auth.AuthClient
}

// Creates a new SalesForceClient
func NewClient(authClient *auth.AuthClient) *SalesForceClient {
	return &SalesForceClient{authClient}
}
