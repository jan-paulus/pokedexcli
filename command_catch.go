package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("Please provide a name or id of a pokemon")
	}

	idOrAreaName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(idOrAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	rand := rand.Intn(pokemon.BaseExperience)

	if rand > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

  fmt.Printf("%s was caught!\n", pokemon.Name)
  cfg.caughtPokemon[pokemon.Name] = pokemon 
	return nil
}
