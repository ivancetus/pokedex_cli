package main

import (
	"github.com/ivancetus/pokedex_cli/config"
	"github.com/ivancetus/pokedex_cli/internal/pokeapi"
	"time"
)

func main() {
	cacheInterval := time.Hour
	cfg := config.Config{
		PokeApiClient: pokeapi.NewClient(cacheInterval),
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

}
