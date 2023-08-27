package commands

import (
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackHelp(cfg *config.Config, s ...string) error {
	commandsMap, commandsSlice := GetCommands()

	fmt.Println("Welcome to the Pokedex_cli help menu!")
	fmt.Println("\nAvailable Commands:")
	for _, key := range commandsSlice {
		fmt.Printf("  %-*s%-*s\n", 24, commandsMap[key].Name, 40, commandsMap[key].Description)
	}
	fmt.Println("")
	return nil
}
