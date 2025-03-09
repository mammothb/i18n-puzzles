package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// infile, err := os.Open("../data/test_input.txt")
	infile, err := os.Open("../data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	counter := make(map[int64]int)
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Text()
		t, err := time.Parse(time.RFC3339, line)
		if err != nil {
			log.Fatal(err)
		}
		unix_time := t.Unix()
		counter[unix_time]++
		if counter[unix_time] >= 4 {
			fmt.Println(t.UTC().Format("2006-01-02T15:04:05-07:00"))
			break
		}
	}
}

