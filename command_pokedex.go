package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("your pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p)
	}
	return nil
}
