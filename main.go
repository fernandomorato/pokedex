package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	cfg := &config{
		client: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	repl(cfg)
}
