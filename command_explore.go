package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	locationArea, err := cfg.client.GetLocationArea(args[0])
	if err != nil {
		return err
	}
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No pokemons in this area :(")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, e := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}
	return nil
}
