package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeCache"
)

type config struct {
	Next  string
	Prev  string
	Cache *pokeCache.Cache
}

func repl() error {
	scanner := bufio.NewScanner(os.Stdin)

	appConfig := config{}
	appConfig.Cache = pokeCache.NewCache(time.Duration(5) * time.Second)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := scanner.Text()

		if len(words) == 0 {
			continue
		}

		params := strings.Split(words, " ")

		if len(params) == 1 {
			params = append(params, "")
		}
		cmd, ok := getCommands()[params[0]]
		if !ok {
			fmt.Printf("Unrecognized command :: %v\n", words)
			continue
		}

		err := cmd.callback(&appConfig, params[1])
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
