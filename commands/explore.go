package commands

import (
	"errors"
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.PokeApiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s\n", locationArea.Name)
	for _, p := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
