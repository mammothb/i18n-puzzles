package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mammothb/i18n-puzzles/helper"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	var lines [][]rune
	var puzzles []string
	var b bytes.Buffer
	transformer := transform.NewWriter(&b, charmap.ISO8859_1.NewEncoder())
	transformer2 := transform.NewWriter(&b,
		transform.Chain(charmap.ISO8859_1.NewEncoder(), charmap.ISO8859_1.NewEncoder()))
	i := 0
	// for line := range helper.ReadLine("data/test_input.txt") {
	for line := range helper.ReadLine("data/input.txt") {
		i++
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, ".") {
			puzzles = append(puzzles, strings.TrimSpace(line))
		} else if i%3 == 0 && i%5 == 0 {
			transformer2.Write([]byte(line))
			lines = append(lines, []rune(b.String()))
		} else if i%3 == 0 || i%5 == 0 {
			transformer.Write([]byte(line))
			lines = append(lines, []rune(b.String()))
		} else {
			lines = append(lines, []rune(line))
		}
		b.Reset()
	}

	result := 0
	for _, puzzle := range puzzles {
		l := utf8.RuneCountInString(puzzle)
		trimmed := []rune(strings.TrimLeft(puzzle, "."))
		idx := l - len(trimmed)
		for i, line := range lines {
			if len(line) != l {
				continue
			}
			if line[idx] == trimmed[0] {
				result += i + 1
			}
		}
	}
	fmt.Println(result)
}
