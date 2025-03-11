package main

import (
	"fmt"

	"github.com/mammothb/i18n-puzzles/helper"
)

func main() {
	result := 0
	i := 0
	// for line := range helper.ReadLine("data/test_input.txt") {
	for line := range helper.ReadLine("data/input.txt") {
		runes := []rune(line)
		if runes[i] == []rune("💩")[0] {
			result++
		}
		i = (i + 2) % len(runes)
	}

	fmt.Println(result)
}
