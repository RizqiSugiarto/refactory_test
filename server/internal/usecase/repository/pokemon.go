package repository

import (
	"database/sql"
	"fmt"

	"github.com/MiniProject/go-oauth2/internal/entity"
)

type Pokemon struct {
	*sql.DB
}

func NewPokeRepo(db *sql.DB) *Pokemon {
	return &Pokemon{db}
}

func (p *Pokemon) InsertPoke(pokeData entity.Pokemon) error {
	sqlStatement := `
	INSERT INTO pokemons (name, url)
	VALUES ($1, $2)`

	_, err := p.DB.Exec(sqlStatement,
		pokeData.Name,
		pokeData.Url,
	)

	if err != nil {
		return fmt.Errorf("error when insert data into database: %s", err)
	}
	return nil
}
