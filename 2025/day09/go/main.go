package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/mammothb/i18n-puzzles/helper"
)

var timeLayouts = []string{"02-01-06", "01-02-06", "06-01-02", "06-02-01"}

func isNineEleven(date time.Time) bool {
	return date.Year() == 2001 && date.Month() == time.September && date.Day() == 11
}

func main() {
	nameToDates := make(map[string][]string)
	// for line := range helper.ReadLine("test_input.txt") {
	for line := range helper.ReadLine("input.txt") {
		parts := strings.Split(line, ": ")
		names := strings.Split(parts[1], ", ")
		for _, name := range names {
			nameToDates[name] = append(nameToDates[name], parts[0])
		}
	}
	result := []string{}
	for name, dates := range nameToDates {
		for _, layout := range timeLayouts {
			ok := true
			parsedDates := []time.Time{}
			for _, date := range dates {
				t, err := time.Parse(layout, date)
				if err != nil {
					ok = false
					break
				}
				parsedDates = append(parsedDates, t)
			}
			if ok {
				for _, date := range parsedDates {
					if isNineEleven(date) {
						result = append(result, name)
					}
				}
			}
		}
	}
	slices.Sort(result)
	fmt.Println(strings.Join(result, " "))
}
