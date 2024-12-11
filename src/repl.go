package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeApi"
	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeCache"
)

type config struct {
	Next               string
	Prev               string
	Cache              *pokeCache.Cache
	CurrentArea        string
	CurrentAreaPokemon []string
	Pokedex            map[string]pokeApi.PokemonResponse
}

func repl() error {
	scanner := bufio.NewScanner(os.Stdin)

	appConfig := config{}
	appConfig.Cache = pokeCache.NewCache(time.Duration(60) * time.Second)
	appConfig.Pokedex = map[string]pokeApi.PokemonResponse{}

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := scanner.Text()

		if len(words) == 0 {
			continue
		}

		params := cleanInput(words)

		if len(params) == 1 {
			params = append(params, "")
		}
		cmd, ok := getCommands()[params[0]]
		if !ok {
			fmt.Printf("Unrecognized command :: %v\n", words)
			continue
		}

		err := cmd.callback(&appConfig, params[1:]...)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {

	words := strings.Split(strings.TrimSpace(text), " ")
	clean_words := []string{}

	for _, word := range words {
		if strings.Contains(word, " ") || word == "" {
			continue
		}

		clean_words = append(clean_words, word)
	}

	return clean_words
}
