package commands

import (
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackPokedex(cfg *config.Config, args ...string) error {
	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	fmt.Println()
	return nil
}
