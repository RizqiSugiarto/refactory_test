package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	"golang.org/x/oauth2/google"
)

type GoogleUrl struct {
	AuthCodeUrl         string
	GoogleUserInfoToken string
}

type Config struct {
	GoogleLoginConfig oauth2.Config
	UrlPostgresDb     string
	GoogleUrl
}

var AppConfig Config

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{os.Getenv("GOOGLE_URL_USERINFO_EMAIL"),
			os.Getenv("GOOGLE_URL_USERINFO_PROFILE")},
		Endpoint: google.Endpoint,
	}

	AppConfig.GoogleUserInfoToken = os.Getenv("GOOGLE_USER_INFO_WITH_TOKEN")
	AppConfig.AuthCodeUrl = os.Getenv("GOOGLE_AUTH_CODE_URL")

	AppConfig.UrlPostgresDb = os.Getenv("DATABASE_URL")

	return &AppConfig
}
