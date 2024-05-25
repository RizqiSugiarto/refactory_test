package usecase

import "github.com/Siddheshk02/go-oauth2/internal/entity"

type (
	UserRepo interface {
		InsertDataUser(userData entity.User) error
	}

	UserService interface {
		GoogleLogin(code string) (string, error)
		GoogleLogins() string
	}

	PokemonRepo interface {
		InsertPoke(pokeData entity.Pokemon) error
	}

	PokemonService interface {
		FetchPokemon() ([]entity.Pokemon, error)
	}
)
