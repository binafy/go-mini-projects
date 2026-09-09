package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Card numbers are 12 to 19 digits long, per ISO/IEC 7812.
const (
	minLength = 12
	maxLength = 19
)

// go run main.go -number="4539 1488 0343 6467"
func main() {
	number := flag.String("number", "", "Card number to validate (reads arguments, then stdin, when empty)")
	flag.Parse()

	numbers, err := readNumbers(*number, flag.Args(), os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if len(numbers) == 0 {
		log.Fatal("No card number given: pass -number, an argument, or pipe numbers into stdin")
	}

	valid := 0
	for _, input := range numbers {
		digits := clean(input)

		if !isValid(digits) {
			fmt.Printf("'%s' is not a valid card number.\n", input)
			continue
		}

		valid++
		network := cardType(digits)
		if network == "" {
			fmt.Printf("'%s' is a valid card number from an unknown network.\n", input)
			continue
		}

		fmt.Printf("'%s' is a valid %s card number.\n", input, network)
	}

	fmt.Printf("\nChecked %d card number(s): %d valid, %d invalid.\n", len(numbers), valid, len(numbers)-valid)
}

// readNumbers collects the card numbers to check: the -number flag, the
// remaining arguments or, when both are empty, one number per line from stdin.
func readNumbers(number string, args []string, stdin io.Reader) ([]string, error) {
	if number != "" {
		return []string{number}, nil
	}

	if len(args) > 0 {
		return args, nil
	}

	numbers := make([]string, 0)
	scanner := bufio.NewScanner(stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			numbers = append(numbers, line)
		}
	}

	return numbers, scanner.Err()
}

// clean strips the spaces and dashes people type card numbers with.
func clean(input string) string {
	return strings.NewReplacer(" ", "", "-", "").Replace(input)
}

// isValid reports whether the number has a plausible length, holds nothing but
// digits and passes the Luhn checksum.
func isValid(digits string) bool {
	if len(digits) < minLength || len(digits) > maxLength {
		return false
	}

	sum := 0
	double := false

	// Luhn: walking from the right, every second digit is doubled, and a
	// doubled result above 9 has 9 subtracted from it. A valid number leaves a
	// total that divides by 10.
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < '0' || digits[i] > '9' {
			return false
		}

		digit := int(digits[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}

// cardType names the network that issued a number, based on the digits it starts with.
func cardType(digits string) string {
	switch {
	case hasPrefix(digits, "4"):
		return "Visa"
	case hasPrefix(digits, "34", "37"):
		return "American Express"
	case hasPrefix(digits, "51", "52", "53", "54", "55"), digits[:4] >= "2221" && digits[:4] <= "2720":
		return "Mastercard"
	case hasPrefix(digits, "6011", "65"):
		return "Discover"
	case hasPrefix(digits, "35"):
		return "JCB"
	case hasPrefix(digits, "30", "36", "38", "39"):
		return "Diners Club"
	case hasPrefix(digits, "62"):
		return "UnionPay"
	}

	return ""
}

// hasPrefix reports whether digits starts with any one of the prefixes.
func hasPrefix(digits string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(digits, prefix) {
			return true
		}
	}

	return false
}
