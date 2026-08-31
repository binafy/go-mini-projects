package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars  = "0123456789"
	symbolChars = "!@#$%^&*()-_=+[]{};:,.?"
)

// go run main.go -length=20 -count=3
func main() {
	length := flag.Int("length", 16, "Length of each password")
	count := flag.Int("count", 1, "Number of passwords to generate")
	lower := flag.Bool("lower", true, "Include lowercase letters")
	upper := flag.Bool("upper", true, "Include uppercase letters")
	digits := flag.Bool("digits", true, "Include digits")
	symbols := flag.Bool("symbols", true, "Include symbols")
	flag.Parse()

	if *count < 1 {
		log.Fatal("Number of passwords must be at least 1")
	}

	classes := selectedClasses(*lower, *upper, *digits, *symbols)
	if len(classes) == 0 {
		log.Fatal("At least one character class must be enabled")
	}
	if *length < len(classes) {
		log.Fatalf("Length must be at least %d to fit every enabled character class", len(classes))
	}

	for i := 0; i < *count; i++ {
		password, err := generatePassword(*length, classes)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(password)
	}
}

// selectedClasses returns the character sets the user asked for.
func selectedClasses(lower, upper, digits, symbols bool) []string {
	classes := make([]string, 0, 4)

	for _, class := range []struct {
		enabled bool
		chars   string
	}{
		{lower, lowerChars},
		{upper, upperChars},
		{digits, digitChars},
		{symbols, symbolChars},
	} {
		if class.enabled {
			classes = append(classes, class.chars)
		}
	}

	return classes
}

// generatePassword builds one password of the given length, guaranteeing at
// least one character from every enabled class. Characters come from
// crypto/rand, so the result is safe to actually use as a password.
func generatePassword(length int, classes []string) (string, error) {
	alphabet := strings.Join(classes, "")
	password := make([]byte, 0, length)

	for _, class := range classes {
		char, err := randomChar(class)
		if err != nil {
			return "", err
		}

		password = append(password, char)
	}

	for len(password) < length {
		char, err := randomChar(alphabet)
		if err != nil {
			return "", err
		}

		password = append(password, char)
	}

	// The guaranteed characters always land at the front, which would leak the
	// class order, so shuffle before handing the password over.
	if err := shuffle(password); err != nil {
		return "", err
	}

	return string(password), nil
}

// randomChar picks one character from chars using a uniform random index.
func randomChar(chars string) (byte, error) {
	index, err := randomInt(len(chars))
	if err != nil {
		return 0, err
	}

	return chars[index], nil
}

// shuffle reorders the password in place with a Fisher-Yates shuffle.
func shuffle(password []byte) error {
	for i := len(password) - 1; i > 0; i-- {
		j, err := randomInt(i + 1)
		if err != nil {
			return err
		}

		password[i], password[j] = password[j], password[i]
	}

	return nil
}

// randomInt returns a uniform random number in [0, max).
func randomInt(max int) (int, error) {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("reading random data: %w", err)
	}

	return int(number.Int64()), nil
}
