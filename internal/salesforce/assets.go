package salesforce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alireza-msv/jet/internal/utils"
)

type AssetsRequest struct {
	Fields []string     `json:"fields,omitempty"`
	Page   PageObject   `json:"pageObject,omitempty"`
	Query  QueryObject  `json:"queryObject,omitempty"`
	Sort   []SortObject `json:"sort,omitempty"`
}

type AssetsResponse struct {
	Count    int            `json:"count"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	Links    map[string]any `json:"links"`
	Items    []AssetItem    `json:"items"`
}

type AssetType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"string"`
}

type AssetItem struct {
	ID                       int64             `json:"id"`
	CustomerKey              string            `json:"customerKey"`
	ContentType              string            `json:"contentType"`
	AssetType                AssetType         `json:"assetType"`
	Data                     any               `json:"data"`
	Version                  int               `json:"version"`
	Locked                   bool              `json:"locked"`
	FileProperties           any               `json:"fileProperties"`
	Name                     string            `json:"name"`
	Description              string            `json:"description"`
	Category                 any               `json:"category"`
	ObjectID                 string            `json:"objectID"`
	Tags                     []any             `json:"tags"`
	Content                  string            `json:"content"`
	Design                   string            `json:"design"`
	SuperContent             string            `json:"superContent"`
	CustomFields             any               `json:"customFields"`
	Views                    any               `json:"views"`
	Blocks                   any               `json:"blocks"`
	MinBlocks                int               `json:"minBlocks"`
	MaxBlocks                int               `json:"maxBlocks"`
	Channels                 any               `json:"channels"`
	AllowedBlocks            []any             `json:"allowedBlocks"`
	Slots                    any               `json:"slots"`
	BusinessUnitAvailability any               `json:"businessUnitAvailability"`
	Template                 any               `json:"template"`
	File                     string            `json:"file"`
	GenerateFrom             string            `json:"generatedFrom"`
	SharingProperties        SharingProperties `json:"sharingProperties"`
	CreatedDate              string            `json:"createdDate"`
	ModifiedDate             string            `json:"modifiedDate"`
}

type SharingProperties struct {
	SharedWith  []string `json:"sharedWith"`
	SharingType string   `json:"sharingType"`
}

func (client *SalesForceClient) QueryAssets(request AssetsRequest) (*AssetsResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	data, err := client.makeRequest(http.MethodPost, utils.AssetsQueryURL, bytes.NewReader(b))

	assetsResponse := AssetsResponse{}
	err = json.Unmarshal(data, &assetsResponse)

	return &assetsResponse, nil
}

func (client *SalesForceClient) makeRequest(method string, path string, body io.Reader) ([]byte, error) {
	accessToken, err := client.authClient.Token()
	if err != nil {
		return nil, err
	}

	reqUrl, err := url.JoinPath(client.authClient.RESTURI(), path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(utils.HttpContentTypeHeader, utils.HttpContentTypeJSONHeader)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBytes))
	}

	return resBytes, nil
}
