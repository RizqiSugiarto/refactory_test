package controllers

import (
	"fmt"

	"github.com/Siddheshk02/go-oauth2/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type PokemonController struct {
	service usecase.PokemonService
}

func NewPoke(service usecase.PokemonService, h *fiber.App) {
	p := &PokemonController{service: service}

	h.Get("/pokemon", p.GetPokemon)
}

func (p *PokemonController) GetPokemon(c *fiber.Ctx) error {
	pokeData, err := p.service.FetchPokemon()

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(pokeData)
}
