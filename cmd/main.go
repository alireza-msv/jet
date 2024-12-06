package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alireza-msv/jet/internal/auth"
	"github.com/alireza-msv/jet/internal/config"
	"github.com/alireza-msv/jet/internal/salesforce"
	"github.com/alireza-msv/jet/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	authClient := auth.NewAuthClient(cfg.Subdomain, auth.ClientOptions{
		AccountID:    cfg.AccountID,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	})

	client := salesforce.NewClient(authClient)

	strg := NewStorage()

	cursor, err := strg.GetAssetsCursor()
	if err != nil {
		log.Fatal("Failed to get the cursor")
	}

	page := 1
	pageSize := 50
	assetsRequest := salesforce.AssetsRequest{
		Page: salesforce.PageObject{Page: page, PageSize: pageSize},
		Query: salesforce.QueryObject{
			LeftOperand: salesforce.QueryOperand{
				Property:       "createdDate",
				SimpleOperator: salesforce.OperandGreaterThan,
				Value:          cursor,
			},
			LogicalOperator: salesforce.LogicalOperandOr,
			RightOperand: salesforce.QueryOperand{
				Property:       "modifedDate",
				SimpleOperator: salesforce.OperandGreaterThan,
				Value:          cursor,
			},
		},
		Sort: []salesforce.SortObject{
			{
				Property:  "createdDate",
				Direction: salesforce.SortDirectionAscending,
			},
			{
				Property:  "modifiedDate",
				Direction: salesforce.SortDirectionAscending,
			},
		},
	}

	for {
		assets, err := client.QueryAssets(assetsRequest)
		if err != nil {
			// TODO: implement a better error handling
			fmt.Print("Error on getting assets::", err)
			break
		}

		fmt.Printf("%d Asset found \n", assets.Count)

		if len(assets.Items) == 0 {
			// No more assets left
			break
		} else {
			strg.SaveAssets(&assets.Items)

			assetsRequest.Page.Page += 1
		}
	}

	err = strg.SaveAssetsCursor(time.Now().UTC().String())
	if err != nil {
		fmt.Println("Error on saving the cursor::", err)
	}
}

// The function is just a simple demonstration the power of using interfaces
// over single purpose structs
func NewStorage() storage.Storage {
	// Based on the config a new instance of Storage (e.g. S3, Local file or Google Cloud Storage)
	// will be created and returned
	return storage.NewLocalStorage()
}
