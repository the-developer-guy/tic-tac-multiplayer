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
	LocalTest     bool
}

func LoadConfig() (*AppConfig, error) {

	ac := AppConfig{
		AdminUser:     os.Getenv("TTTSERVER_ADMIN_USER"),
		AdminPassword: os.Getenv("TTTSERVER_ADMIN_PASS"),
		AdminToken:    os.Getenv("TTTSERVER_ADMIN_TOKEN"),
		LocalTest:     false,
	}

	lt := os.Getenv("TTTSERVER_LOCAL_TEST")
	if lt != "" {
		fmt.Println("Loopback/test mode set")

		ac.LocalTest = true

		if ac.AdminUser == "" {
			ac.AdminUser = "admin"
		}
		if ac.AdminPassword == "" {
			ac.AdminPassword = "admin"
		}
		if ac.AdminToken == "" {
			ac.AdminToken = "admin"
		}

		return &ac, nil
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
