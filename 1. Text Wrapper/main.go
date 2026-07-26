package main

import (
	"flag"
	"fmt"
	"strings"
)

// go run main.go -text="Im Milwad Khosravi who make a lot of packages and tools for developers, enjoy it!" -width=10
func main() {
	text := flag.String("text", "", "Text to wrap")
	width := flag.Int("width", 5, "Maximum line width")
	flag.Parse()

	if *text == "" {
		fmt.Println("Text is required")
		return
	}

	words := strings.Fields(*text)

	var lines []string
	var currentLine string

	for _, word := range words {
		// First word on the line
		if currentLine == "" {
			currentLine = word
			continue
		}

		// Check if adding the next word exceeds the width
		if len(currentLine)+1+len(word) <= *width {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	// Add the last line
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	fmt.Println(strings.Join(lines, "\n"))
}
