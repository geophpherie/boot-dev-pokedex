package main

import (
	"fmt"
)

func main() {
	err := repl()
	if err != nil {
		fmt.Printf("unhandled error in repl :: %v", err)
	}
}
