package main

import (
	"bufio"
	"fmt"
	"github.com/ivancetus/pokedex_cli/commands"
	"github.com/ivancetus/pokedex_cli/config"
	"os"
	"strings"
)

func startRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		allCommands := commands.GetCommands()
		command, ok := allCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := command.Callback(cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
