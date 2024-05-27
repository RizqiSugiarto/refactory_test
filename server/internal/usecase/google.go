package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MiniProject/go-oauth2/internal/config"
	"github.com/MiniProject/go-oauth2/internal/entity"
)

type googleService struct {
	cfg  *config.Config
	repo UserRepo
}

func New(cfg *config.Config, repo UserRepo) *googleService {
	return &googleService{
		cfg:  cfg,
		repo: repo,
	}
}

func (g *googleService) GoogleLogin(code string) (string, error) {
	token, err := g.cfg.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("failed exchange token code: %w", err)
	}

	resp, err := http.Get(g.cfg.GoogleUserInfoToken + token.AccessToken)
	if err != nil {
		return "", fmt.Errorf("failed user data fetch: %w", err)
	}
	dataUser, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse userData: %w", err)
	}

	var fetchApiUser entity.User

	if err := json.Unmarshal(dataUser, &fetchApiUser); err != nil {
		return "", fmt.Errorf("failed to parse fetchData into struct: %w", err)
	}

	if err := g.repo.InsertDataUser(fetchApiUser); err != nil {
		return "", err
	}
	return string(dataUser), nil
}

func (g *googleService) GoogleLogins() string {
	return g.cfg.GoogleLoginConfig.AuthCodeURL(g.cfg.AuthCodeUrl)
}
