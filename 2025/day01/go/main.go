package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	// infile, err := os.Open("../data/test_input.txt")
	infile, err := os.Open("../data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	result := 0

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Text()
		num_bytes := len(line)
		num_chars := utf8.RuneCountInString(line)
		if num_bytes <= 160 && num_chars <= 140 {
			result += 13
		} else if num_bytes <= 160 {
			result += 11
		} else if num_chars <= 140 {
			result += 7
		}
	}
	fmt.Println(result)
}
