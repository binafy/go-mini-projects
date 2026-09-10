package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const description = "Golang, also known as Go, is an open-source programming language developed by Google. It is designed to be simple, fast, and efficient. Go has built-in support for concurrency, which makes it a popular choice for web servers, cloud applications, APIs, and distributed systems."

func main() {
	text := flag.String("text", "", "text to search")
	flag.Parse()

	if *text == "" {
		log.Fatal("Must provide a text to search")
	}

	pattern := `(?i)\b` + regexp.QuoteMeta(*text) + `\b`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringIndex(description, -1)

	for index, match := range matches {
		fmt.Printf("%d Index: [%d:%d]\n", index, match[0], match[1])
	}

	fmt.Printf("Text: %s\n", *text)
	fmt.Printf("Repeat count: %d\n", len(matches))
}
