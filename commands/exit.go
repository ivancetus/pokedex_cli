package commands

import (
	"github.com/ivancetus/pokedex_cli/config"
	"os"
)

func callbackExit(cfg *config.Config, s ...string) error {
	os.Exit(0)
	return nil
}
