package internal

import (
	"errors"
	"fmt"
	"os"
)

type AppConfig struct {
	AdminUser     string
	AdminPassword string
	AdminToken    string
	Standalone    bool
}

func LoadConfig() (*AppConfig, error) {

	ac := AppConfig{
		AdminUser:     os.Getenv("TTTSERVER_ADMIN_USER"),
		AdminPassword: os.Getenv("TTTSERVER_ADMIN_PASS"),
		AdminToken:    os.Getenv("TTTSERVER_ADMIN_TOKEN"),
		Standalone:    false,
	}

	standalone := os.Getenv("TTTSERVER_STANDALONE")
	if standalone != "" {
		ac.Standalone = true
		fmt.Println("Standalone mode set")
	}

	if ac.AdminUser == "" {
		return nil, errors.New("missing admin username from config")
	}
	if ac.AdminPassword == "" {
		return nil, errors.New("missing admin password from config")
	}
	if ac.AdminToken == "" {
		return nil, errors.New("missing admin token from config")
	}

	return &ac, nil
}
