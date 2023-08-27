package commands

import (
	"fmt"
	"github.com/ivancetus/pokedex_cli/config"
)

func callbackHelp(cfg *config.Config) error {
	allCommands := GetCommands()

	fmt.Println("Welcome to the Pokedex_cli help menu!")
	fmt.Println("\nAvailable Commands:")
	for _, cmd := range allCommands {
		fmt.Printf("  %-*s%-*s\n", len(cmd.Name)+2, cmd.Name, len(cmd.Description)+2, cmd.Description)
	}
	fmt.Println("")
	return nil
}
