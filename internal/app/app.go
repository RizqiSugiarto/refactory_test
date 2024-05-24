package app

import (
	"log"

	"github.com/Siddheshk02/go-oauth2/internal/config"
	"github.com/Siddheshk02/go-oauth2/internal/controllers"
	"github.com/Siddheshk02/go-oauth2/internal/usecase"
	"github.com/Siddheshk02/go-oauth2/internal/usecase/repository"
	"github.com/Siddheshk02/go-oauth2/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	cfg := config.LoadConfig()
	RunMigrate(cfg.UrlPostgresDb)
	pg, err := db.ConnectPostgres(cfg.UrlPostgresDb)
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}
	repo := repository.New(pg)
	uc := usecase.New(cfg, repo)
	controllers.New(uc, app)
	app.Listen(":8080")
}