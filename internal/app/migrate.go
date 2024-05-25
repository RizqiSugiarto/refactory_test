package app

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrate(url string) {
	// url += "?sslmode=disable"
	fmt.Println(url)
	m, err := migrate.New(
		"file://migrations",
		url)

	if err != nil {
		log.Fatal("app - runMigrate", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("app - runMigrate", err)
	}
}
