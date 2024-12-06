package config

import (
	"fmt"
	"os"

	"github.com/alireza-msv/jet/internal/utils"
)

type Config struct {
	AccountID    string
	ClientID     string
	ClientSecret string
	Subdomain    string
	Schedule     string
}

func LoadConfig() (*Config, error) {
	accountID, err := readEnvVariable("ACCOUNT_ID", true)
	if err != nil {
		return nil, err
	}

	clientID, err := readEnvVariable("CLIENT_ID", true)
	if err != nil {
		return nil, err
	}

	clientSecret, err := readEnvVariable("CLIENT_SECRET", true)
	if err != nil {
		return nil, err
	}

	subdomain, err := readEnvVariable("SUBDOMAIN", true)
	if err != nil {
		return nil, err
	}

	schedule, err := readEnvVariable("SCHEDULE", false)
	if schedule != "" {
		schedule = utils.DefaultSchedule
	}

	return &Config{
		AccountID:    accountID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Subdomain:    subdomain,
		Schedule:     schedule,
	}, nil
}

func readEnvVariable(name string, required bool) (string, error) {
	val := os.Getenv(name)
	if required && val == "" {
		return "", fmt.Errorf("Env variable not present: %s", name)
	}

	return val, nil
}
