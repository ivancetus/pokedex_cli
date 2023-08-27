package main

import (
	"github.com/ivancetus/pokedex_cli/config"
	"github.com/ivancetus/pokedex_cli/internal/pokeapi"
)

func main() {
	cfg := config.Config{
		PokeApiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)

}
