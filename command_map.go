package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	res, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationUrl)

	if err != nil {
		return err
	}

	cfg.previousLocationUrl = res.Previous
	cfg.nextLocationUrl = res.Next

	for _, element := range res.Results {
		fmt.Println(element.Name)
	}
	return nil
}

func commandMapB(cfg *config, args []string) error {
	if cfg.previousLocationUrl == nil {
		return errors.New("You are on the first page")
	}

	res, err := cfg.pokeapiClient.GetLocations(cfg.previousLocationUrl)

	if err != nil {
		return err
	}

	cfg.previousLocationUrl = res.Previous
	cfg.nextLocationUrl = res.Next

	for _, element := range res.Results {
		fmt.Println(element.Name)
	}
	return nil
}
