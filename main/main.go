package main

import (
	ascii "ascii/pkg"
	"fmt"
	"os"
)

func main() {
	// Handle args error
	err := ascii.ArgsErrors()
	if err {
		return
	}

	// Convert input to art and print result
	art := ascii.AsciiArtFS(os.Args[1])
	fmt.Print(art)
}
