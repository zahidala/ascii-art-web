package ascii

import (
	"os"
	"strings"
)

// Algorithm to find first line of each character
func getFirstLine(char rune) int {
	var firstLine int
	diff := int(char) - 32
	firstLine = diff*9 + 2
	return firstLine
}

// Function to get input as art
func AsciiArtFS(arg string, fileName string) (string, bool, int) {
	var art string
	// Split args by each new line
	phrases := strings.Split(arg, "\r\n")

	// Print ascii for each phrase that was separated by a new line
	for i, phrase := range phrases {
		if len(phrase) == 0 {
			if i < len(phrases)-1 { // Stop printing extra line
				art += "\n"
			}
			continue
		}
		var firstLines []int
		// Save all first lines of each character into a slice
		for _, char := range phrase {
			if char < 32 || char > 126 {
				return "Character out of range\n", true, 400
			} else {
				firstLine := getFirstLine(char)
				firstLines = append(firstLines, firstLine)
			}
		}

		// Read the file

		file, err := os.ReadFile("fonts/" + fileName + ".txt")

		if err != nil {
			return "Could not read font file\n", true, 500
		}

		// Split the data into a slice of strings, one for each line.
		var lines []string

		lines = strings.Split(string(file), "\n")

		// /r is used to take care of carriage return character as only thinkertoy font file uses it.
		if fileName == "thinkertoy" {
			lines = strings.Split(string(file), "\r\n")
		}

		for i := 1; i < 9; i++ {
			for j, line := range firstLines {
				filteredLines := lines[line-1]
				art += filteredLines
				firstLines[j]++
			}
			art += "\n"
		}
	}
	return art, false, 200
}
