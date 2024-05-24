package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgres(url string) (*sql.DB, error) {
	result, err := sql.Open("postgres", url)

	if err != nil {
		return &sql.DB{}, fmt.Errorf("failed to open connectio %s", err)
	}

	return result, nil
}
