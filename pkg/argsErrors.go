package ascii

import (
	"fmt"
	"os"
)

func ArgsErrors() bool {
	// Error handling os arguments
	// Run only if 1 arg
	if len(os.Args) < 2 {
		fmt.Println("ERROR: No arguments to print")
		return true
	} else if len(os.Args) > 3 {
		fmt.Println("ERROR: Place all text in a single argument")
		return true
	} else if len(os.Args[1]) < 1 {
		return true
	}
	return false
}
