package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
)

func repl(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := cleanInput(sc.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		if value, ok := getCommands()[command]; ok {
			err := value.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.Trim(text, " ")))
}

type config struct {
	client   pokeapi.Client
	previous *string
	next     *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspects one of your caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows your caught pokemons",
			callback:    commandPokedex,
		},
	}
}
