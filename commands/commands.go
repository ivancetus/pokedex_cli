package commands

import "github.com/ivancetus/pokedex_cli/config"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() (map[string]cliCommand, []string) {
	commands := map[string]cliCommand{
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
		"explore": {
			Name:        "explore {location_are}",
			Description: "List the pokemon in the location area",
			Callback:    callbackExplore,
		},
		"catch": {
			Name:        "catch {pokemon_name}",
			Description: "Attempt to catch pokemon and add to pokedex",
			Callback:    callbackCatch,
		},
		"inspect": {
			Name:        "inspect {pokemon_name}",
			Description: "View information about a caught pokemon",
			Callback:    callbackInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "View all caught pokemons",
			Callback:    callbackPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit Pokedex_cli",
			Callback:    callbackExit,
		},
	}
	orderedKeys := []string{
		"help", "map", "mapb", "explore", "catch", "inspect", "pokedex", "exit",
	}
	//for key := range commands {
	//	orderedKeys = append(orderedKeys, key)
	//}
	return commands, orderedKeys
}
