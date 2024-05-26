package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	var cobak []entity.Pokemon
	data, _ := json.Marshal(respMap.Result)

	if err := json.Unmarshal(data, &cobak); err != nil {
		return []entity.Pokemon{}, fmt.Errorf("failed parse resp: %w", err)
	}

	for i, pokeDet := range cobak {

		resps, _ := http.Get(pokeDet.Url)

		body, err := io.ReadAll(resps.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		var datas map[string]interface{}
		if err := json.Unmarshal(body, &datas); err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}

		pokeDet.Url = datas["sprites"].(map[string]interface{})["front_default"].(string)

		if err := p.repo.InsertPoke(pokeDet); err != nil {
			return nil, fmt.Errorf("error inserting pokemon: %w", err)
		}
		respMap.Result[i].Url = pokeDet.Url

	}

	return respMap.Result, nil
}
