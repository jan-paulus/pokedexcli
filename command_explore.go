package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("Please provide a name or id of a location")
	}

	idOrAreaName := args[0]

	res, err := cfg.pokeapiClient.GetLocationDetails(idOrAreaName)

	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", res.Location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range res.PokemonEncounters {
		fmt.Printf("  - %s\n", enc.Pokemon.Name)
	}

	return nil
}
