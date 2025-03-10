package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parseUnix(format string, fields []string) int64 {
	tz, err := time.LoadLocation(fields[1])
	if err != nil {
		log.Fatal(err)
	}
	t, err := time.ParseInLocation(format, strings.Join(fields[2:], " "), tz)
	if err != nil {
		log.Fatal(err)
	}
	return t.Unix()
}

func main() {
	// infile, err := os.Open("data/test_input.txt")
	infile, err := os.Open("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	time_format := "Jan 02, 2006, 15:04"
	start_ts := int64(0)
	result := int64(0)

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		if fields[0] == "Departure:" {
			start_ts = parseUnix(time_format, fields)
		} else {
			duration := parseUnix(time_format, fields) - start_ts
			result += duration / 60
		}
	}
	fmt.Println(result)
}
