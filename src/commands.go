package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeApi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
			description: "Show the next location areas to explore",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous location areas to explore",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the given location area for pokemon",
			callback:    commandExplore,
		},
	}
}

func commandMap(config *config, param string) error {
	resp, err := pokeApi.GetLocationAreas(config.Next, config.Cache)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Prev = resp.Previous

	for _, location := range resp.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func commandMapb(config *config, param string) error {
	if config.Prev == "" {
		return errors.New("we're already pulled over, we can't pull over previous any further")
	}

	resp, err := pokeApi.GetLocationAreas(config.Prev, config.Cache)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Prev = resp.Previous

	for _, location := range resp.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func commandExit(config *config, param string) error {
	os.Exit(0)
	return nil
}

func commandHelp(config *config, param string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v:\t%v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExplore(config *config, param string) error {
	if param == "" {
		fmt.Print("No area to explore!")
		return nil
	}
	fmt.Printf("Exploring %v...\n", param)
	resp, err := pokeApi.GetLocationAreaDetail(param, config.Cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		name := encounter.Pokemon.Name
		fmt.Printf(" - %v\n", name)
	}
	return nil
}
