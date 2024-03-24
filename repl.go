package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jan-paulus/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		commandName := input[0]
		args := input[1:]

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show a list of 20 locations, each subsequent call will display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show a list of the previous 20 locations, each subsequent call will display the previous 20 locations, and so on.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of all the Pokemon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a given Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon you caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show a list of all Pokemon you've caught",
			callback:    commandPokedex,
		},
  }
}
