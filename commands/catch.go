package commands

import (
	"errors"
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
	"math/rand"
)

func callbackCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.PokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randNumb := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("\nBaseExperience: %v, RandomNumber: %v, CheckIfGreaterThan: %v",
		pokemon.BaseExperience, randNumb, threshold)
	if randNumb > threshold {
		return fmt.Errorf("\nfailed to catch %s", pokemon.Name)
	}
	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("\n%s was caught!", pokemon.Name)
	fmt.Println()
	return nil
}
