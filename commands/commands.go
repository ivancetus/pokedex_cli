package commands

import "github.com/ivancetus/pokedex_cli/config"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Prints the help menu",
			Callback:    callbackHelp,
		},
		"map": {
			Name:        "map",
			Description: "List the first/next 20 location areas",
			Callback:    callbackMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List the previous 20 location areas",
			Callback:    callbackMapBack,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit Pokedex_cli",
			Callback:    callbackExit,
		},
	}
}
