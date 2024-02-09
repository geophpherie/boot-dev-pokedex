package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("starting")
	fmt.Println("-----")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		fmt.Println(text)
	}
}
