package repository

import (
	"database/sql"
	"fmt"

	"github.com/Siddheshk02/go-oauth2/internal/entity"
)

type UserRepository struct {
	*sql.DB
}

func New(params *sql.DB) *UserRepository {
	return &UserRepository{params}
}

func (u *UserRepository) InsertDataUser(userData entity.User) error {
	sqlStatement := `
	INSERT INTO users (id, email, verified_email, name, given_name, family_picture, picture, locale)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := u.DB.Exec(sqlStatement, userData.Id,
		userData.Email,
		userData.Verified_email,
		userData.Name,
		userData.Given_name,
		userData.Family_picture,
		userData.Picture,
		userData.Locale,
	)

	if err != nil {
		return fmt.Errorf("error when insert data into database: %s", err)
	}
	return nil
}
