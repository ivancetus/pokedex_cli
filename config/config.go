package config

import "github.com/ivancetus/pokedex_cli/internal/pokeapi"

type Config struct {
	PokeApiClient    pokeapi.Client
	NextLocationArea *string
	PrevLocationArea *string
}
