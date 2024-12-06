package storage

import "github.com/alireza-msv/jet/internal/salesforce"

type LocalStorage struct {
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (ls *LocalStorage) SaveAssets(items *[]salesforce.AssetItem) error {
	// Save the assets on the disk
	return nil
}

func (ls *LocalStorage) SaveAssetsCursor(cursor string) error {
	// Save the cursor on the disk
	return nil
}

func (ls *LocalStorage) GetAssetsCursor() (string, error) {
	// Read the cursor from the disk and return it
	return "", nil
}
