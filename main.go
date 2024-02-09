package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := scanner.Text()

		cmd, ok := getCommands()[words]
		if !ok {
			fmt.Printf("Unrecognized command :: %v\n", words)
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
