package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// infile, err := os.Open("data/test_input.txt")
	infile, err := os.Open("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	result := 0

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Text()
		n_rune := utf8.RuneCountInString(line)
		if n_rune < 4 || n_rune > 12 {
			continue
		}
		has_digit := false
		has_upper := false
		has_lower := false
		has_non_ascii := false
		for i, w := 0, 0; i < len(line); i += w {
			r, width := utf8.DecodeRuneInString(line[i:])
			has_digit = has_digit || unicode.IsDigit(r)
			has_upper = has_upper || unicode.IsUpper(r)
			has_lower = has_lower || unicode.IsLower(r)
			has_non_ascii = has_non_ascii || r > unicode.MaxASCII
			w = width
		}
		if has_digit && has_upper && has_lower && has_non_ascii {
			result++
		}
	}
	fmt.Println(result)
}
