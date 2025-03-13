package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mammothb/i18n-puzzles/helper"
)

func deduceOffset(dt time.Time, zones []string) time.Time {
	_, offset := dt.Zone()
	naive := dt.UTC()
	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			log.Fatal(err)
		}
		_, new_offset := naive.In(loc).Zone()
		if new_offset == offset {
			naive = naive.In(loc)
			break
		}
	}
	return naive
}

func main() {
	zones := []string{"America/Halifax", "America/Santiago"}
	result := 0
	i := 0
	for line := range helper.ReadLine("input.txt") {
		i++
		parts := strings.Fields(line)
		dt, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			log.Fatal(err)
		}
		dt = deduceOffset(dt, zones)

		add, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		sub, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		adj_dt := dt.Add(time.Minute * time.Duration(add-sub))

		result += i * adj_dt.Hour()
	}
	fmt.Println(result)
}
