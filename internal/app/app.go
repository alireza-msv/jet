package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/alireza-msv/jet/internal/auth"
	"github.com/alireza-msv/jet/internal/config"
	"github.com/alireza-msv/jet/internal/salesforce"
	"github.com/alireza-msv/jet/internal/storage"
)

type App struct {
	authClient *auth.AuthClient
	client     *salesforce.SalesForceClient
	storage    storage.Storage
}

func NewApp(cfg *config.Config, storage storage.Storage) *App {
	authClient := auth.NewAuthClient(cfg.Subdomain, auth.ClientOptions{
		AccountID:    cfg.AccountID,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	})

	client := salesforce.NewClient(authClient)

	return &App{
		authClient: authClient,
		client:     client,
		storage:    storage,
	}
}

func (a *App) Start() error {
	cursor, err := a.storage.GetAssetsCursor()
	if err != nil {
		return errors.New("Failed to get the cursor")
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
		assets, err := a.client.QueryAssets(assetsRequest)
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
			a.storage.SaveAssets(&assets.Items)

			assetsRequest.Page.Page += 1
		}
	}

	err = a.storage.SaveAssetsCursor(time.Now().UTC().String())
	if err != nil {
		return fmt.Errorf("Error on saving the cursor::", err)
	}

	return nil
}
