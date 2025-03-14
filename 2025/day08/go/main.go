package main

import (
	"fmt"
	"log"
	"slices"
	"unicode"
	"unicode/utf8"

	"github.com/mammothb/i18n-puzzles/helper"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func isVowel(r rune) bool {
	return slices.Contains(vowels, r)
}

func normalize(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {
	result := 0
	for line := range helper.ReadLine("input.txt") {
		l := utf8.RuneCountInString(line)
		if l < 4 || l > 12 {
			continue
		}
		line, err := normalize(line)
		if err != nil {
			log.Fatal(err)
		}
		has_digit := false
		has_vowel := false
		has_consonant := false
		has_repeat := false
		seen := map[rune]struct{}{}
		for _, ch := range line {
			has_digit = has_digit || unicode.IsDigit(ch)
			if !unicode.IsLetter(ch) {
				continue
			}
			ch = unicode.ToLower(ch)
			if isVowel(ch) {
				has_vowel = true
			} else {
				has_consonant = true
			}
			if _, ok := seen[ch]; ok {
				has_repeat = true
				break
			}
			seen[ch] = struct{}{}
		}
		if has_digit && has_vowel && has_consonant && !has_repeat {
			result++
		}
	}
	fmt.Println(result)
}
