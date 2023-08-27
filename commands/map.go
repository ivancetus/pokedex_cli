package commands

import (
	"errors"
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackMap(cfg *config.Config) error {
	res, err := cfg.PokeApiClient.LocationAreaList(cfg.NextLocationArea)
	if err != nil {
		return err
	}
	cfg.NextLocationArea = res.Next
	cfg.PrevLocationArea = res.Previous

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println()
	return nil
}

func callbackMapBack(cfg *config.Config) error {
	if cfg.PrevLocationArea == nil {
		return errors.New("you're on thr first page")
	}
	res, err := cfg.PokeApiClient.LocationAreaList(cfg.PrevLocationArea)
	if err != nil {
		return err
	}
	cfg.NextLocationArea = res.Next
	cfg.PrevLocationArea = res.Previous

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println()
	return nil
}
