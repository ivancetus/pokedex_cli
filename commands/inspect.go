package commands

import (
	"errors"
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}
	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("height: %v\n", pokemon.Height)
	fmt.Printf("weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
