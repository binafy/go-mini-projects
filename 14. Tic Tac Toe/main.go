package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const empty = ' '

// The eight lines that win the game, as board indices.
var lines = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
	{0, 4, 8}, {2, 4, 6}, // diagonals
}

type board [9]rune

func newBoard() board {
	var b board
	for i := range b {
		b[i] = empty
	}

	return b
}

// Empty cells show the number you type to play them.
func (b board) String() string {
	var sb strings.Builder
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 3; col++ {
			if cell := b[row+col]; cell == empty {
				fmt.Fprintf(&sb, " %d ", row+col+1)
			} else {
				fmt.Fprintf(&sb, " %c ", cell)
			}
			if col < 2 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
		if row < 6 {
			sb.WriteString("---+---+---\n")
		}
	}

	return sb.String()
}

func (b board) winner() rune {
	for _, l := range lines {
		if b[l[0]] != empty && b[l[0]] == b[l[1]] && b[l[1]] == b[l[2]] {
			return b[l[0]]
		}
	}

	return empty
}

func (b board) full() bool {
	for _, cell := range b {
		if cell == empty {
			return false
		}
	}

	return true
}

// parseCell turns typed input into a board index, or explains why it cannot.
func parseCell(input string, b board) (int, error) {
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("%q is not a number", input)
	}
	if n < 1 || n > 9 {
		return 0, fmt.Errorf("cell %d is off the board, pick 1-9", n)
	}
	if b[n-1] != empty {
		return 0, fmt.Errorf("cell %d is already taken by %c", n, b[n-1])
	}

	return n - 1, nil
}

func main() {
	xName := flag.String("x", "Player X", "Name of the X player")
	oName := flag.String("o", "Player O", "Name of the O player")
	first := flag.String("first", "x", "Which mark moves first: x or o")
	flag.Parse()

	mark := 'X'
	switch strings.ToLower(*first) {
	case "x":
	case "o":
		mark = 'O'
	default:
		log.Fatalf("-first must be 'x' or 'o', got %q", *first)
	}

	name := map[rune]string{'X': *xName, 'O': *oName}
	b := newBoard()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(b)
		fmt.Printf("%s (%c), pick a cell 1-9: ", name[mark], mark)

		if !scanner.Scan() {
			fmt.Println("\nGame abandoned.")
			return
		}

		// An invalid move must not cost the player their turn.
		cell, err := parseCell(scanner.Text(), b)
		if err != nil {
			fmt.Printf("Invalid move: %v\n\n", err)
			continue
		}

		b[cell] = mark

		if w := b.winner(); w != empty {
			fmt.Println(b)
			fmt.Printf("%s (%c) wins!\n", name[w], w)
			return
		}

		if b.full() {
			fmt.Println(b)
			fmt.Println("It's a draw.")
			return
		}

		if mark == 'X' {
			mark = 'O'
		} else {
			mark = 'X'
		}
	}
}
