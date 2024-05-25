package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Siddheshk02/go-oauth2/internal/entity"
)

type PokemonUseCase struct {
	repo PokemonRepo
}

type PokeResp struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous any              `json:"previous"`
	Result   []entity.Pokemon `json:"results"`
}

func NewPokemon(repo PokemonRepo) *PokemonUseCase {
	return &PokemonUseCase{
		repo: repo,
	}
}

func (p *PokemonUseCase) FetchPokemon() ([]entity.Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")

	if err != nil {
		return []entity.Pokemon{}, fmt.Errorf("error Fetch API: %w", err)
	}

	dataPoke, err := io.ReadAll(resp.Body)
	if err != nil {
		return []entity.Pokemon{}, fmt.Errorf("failed to parse pokeData: %w", err)
	}

	var respMap PokeResp

	if err := json.Unmarshal(dataPoke, &respMap); err != nil {
		return []entity.Pokemon{}, fmt.Errorf("failed parse resp: %w", err)
	}

	// fmt.Println(respMap, "GINI")

	for _, poke := range respMap.Result {
		err := p.repo.InsertPoke(poke)

		if err != nil {
			if err := p.repo.InsertPoke(poke); err != nil {
				return nil, fmt.Errorf("error inserting pokemon: %w", err)
			}
		}
	}

	return respMap.Result, nil
}
