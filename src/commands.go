package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeApi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"catch": {
			name:        "catch",
			description: "Attempts to catch the given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Learn more about a pokemon you've imprisoned",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all your pokemon",
			callback:    commandPokedex,
		},
	}
}

func commandMap(config *config, params ...string) error {
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

func commandMapb(config *config, params ...string) error {
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

func commandExit(config *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, params ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v:\t%v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExplore(config *config, params ...string) error {
	if len(params) != 1 {
		fmt.Println("Invalid area")
		return nil
	}
	area := params[0]
	if area == "" {
		fmt.Println("No area to explore!")
		return nil
	}
	fmt.Printf("Exploring %v...\n", area)
	resp, err := pokeApi.GetLocationAreaDetail(area, config.Cache)
	if err != nil {
		return err
	}

	// reset area pokemon when we explore
	config.CurrentAreaPokemon = []string{}
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		name := encounter.Pokemon.Name
		config.CurrentAreaPokemon = append(config.CurrentAreaPokemon, name)
		fmt.Printf(" - %v\n", name)
	}

	config.CurrentArea = area
	return nil
}

func commandCatch(config *config, params ...string) error {
	pokemonToCatch := params[0]
	if config.CurrentArea == "" {
		fmt.Println("Can't catch anything til you go somewhere.")
		return nil
	}

	if !slices.Contains(config.CurrentAreaPokemon, pokemonToCatch) {
		fmt.Println("Pokemon is not found in this area.")
		return nil
	}

	if _, ok := config.Pokedex[pokemonToCatch]; ok {
		fmt.Println("You already have this pokemon")
		return nil
	}
	resp, err := pokeApi.GetPokemonDetail(pokemonToCatch, config.Cache)
	if err != nil {
		return err
	}

	// calcuate odds of catching
	// resp.BaseExperience
	randCatchProb := rand.Float32()
	experienceMultiplier := 1 - resp.BaseExperience/1000

	catchProb := randCatchProb * float32(experienceMultiplier)

	// attempt catch
	fmt.Printf("Throwing pokeball at %v...\n", pokemonToCatch)

	// add some suspense
	time.Sleep(500 * time.Millisecond)

	// if caught, add to pokedex in config
	if catchProb >= 0.5 {
		fmt.Printf("%v was caught!\n", pokemonToCatch)
		config.Pokedex[pokemonToCatch] = resp
	} else {
		fmt.Printf("%v escaped!\n", pokemonToCatch)
	}
	return nil
}

func commandInspect(config *config, params ...string) error {
	pokemonToInspect := params[0]
	pokemonData, ok := config.Pokedex[pokemonToInspect]

	if !ok {
		fmt.Println("You need to catch this pokemon first!")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemonData.Name)
	fmt.Printf("Height: %v\n", pokemonData.Height)
	fmt.Printf("Weight: %v\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemonData.Stats {
		fmt.Printf(" - %v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemonData.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil

}

func commandPokedex(config *config, params ...string) error {
	fmt.Println("Your pokedex: ")

	for k := range config.Pokedex {
		fmt.Printf(" - %v", k)
	}

	fmt.Println()
	return nil
}
