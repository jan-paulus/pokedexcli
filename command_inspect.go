package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("Please provide the name of the Pokemon you want to inspect")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New("You have not caught that pokemon")
	}
  
  fmt.Printf("Name: %s\n", pokemon.Name)
  fmt.Printf("Height: %s\n", strconv.Itoa(pokemon.Height))
  fmt.Printf("Weight: %s\n", strconv.Itoa(pokemon.Weight))

  fmt.Println("Stats:")
  for _, v := range pokemon.Stats {
    fmt.Printf("  - %s: %s\n", v.Stat.Name, strconv.Itoa(v.BaseStat))
  }


  fmt.Println("Types:")
  for _, v := range pokemon.Types {
    fmt.Printf("  - %s\n", v.Type.Name)
  }

  return nil
}
