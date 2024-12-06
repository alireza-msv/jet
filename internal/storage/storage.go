package storage

import "github.com/alireza-msv/jet/internal/salesforce"

type Storage interface {
	SaveAssets(*[]salesforce.AssetItem) error
	SaveAssetsCursor(string) error
	GetAssetsCursor() (string, error)
}
