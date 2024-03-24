package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You have not caught a Pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
