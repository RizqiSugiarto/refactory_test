package app

import (
	"log"

	"github.com/MiniProject/go-oauth2/internal/config"
	"github.com/MiniProject/go-oauth2/internal/controllers"
	"github.com/MiniProject/go-oauth2/internal/usecase"
	"github.com/MiniProject/go-oauth2/internal/usecase/repository"
	"github.com/MiniProject/go-oauth2/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run() {
	app := fiber.New()
	app.Use(cors.New())

	cfg := config.LoadConfig()
	RunMigrate(cfg.UrlPostgresDb)
	pg, err := db.ConnectPostgres(cfg.UrlPostgresDb)
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}
	repo := repository.New(pg)
	pokeRepo := repository.NewPokeRepo(pg)
	uc := usecase.New(cfg, repo)
	pokeUc := usecase.NewPokemon(pokeRepo)
	controllers.New(uc, cfg, app)
	controllers.NewPoke(pokeUc, app)
	app.Listen(":8080")
}
