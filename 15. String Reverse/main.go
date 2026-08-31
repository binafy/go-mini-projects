package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// go run main.go -text "Milwad Khosravi"
func main() {
	text := flag.String("text", "", "Text to reverse (reads from stdin when empty)")
	words := flag.Bool("words", false, "Reverse the order of the words instead of the characters")
	flag.Parse()

	input, err := readInput(*text, flag.Args(), os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		if *words {
			fmt.Fprintln(writer, reverseWords(line))
			continue
		}

		fmt.Fprintln(writer, reverse(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

// readInput picks the text to reverse: the -text flag, the remaining arguments
// or, when both are empty, everything piped into stdin.
func readInput(text string, args []string, stdin io.Reader) (string, error) {
	if text != "" {
		return text, nil
	}

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	content, err := io.ReadAll(stdin)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// reverse flips a string character by character. Runes are grouped with the
// combining marks that follow them, so "café" stays "éfac" even when the
// accent is stored as a separate mark.
func reverse(text string) string {
	clusters := make([]string, 0, len(text))
	current := make([]rune, 0, 2)

	for _, char := range text {
		if unicode.Is(unicode.Mn, char) && len(current) > 0 {
			current = append(current, char)
			continue
		}

		if len(current) > 0 {
			clusters = append(clusters, string(current))
		}
		current = append(current[:0:0], char)
	}

	if len(current) > 0 {
		clusters = append(clusters, string(current))
	}

	var builder strings.Builder
	builder.Grow(len(text))
	for i := len(clusters) - 1; i >= 0; i-- {
		builder.WriteString(clusters[i])
	}

	return builder.String()
}

// reverseWords flips the order of the words but keeps every word readable.
func reverseWords(text string) string {
	words := strings.Fields(text)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
