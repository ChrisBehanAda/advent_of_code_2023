package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Solve is the entry point for running the solutions.
func Solve() {
	fmt.Println("Day 1:")
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, l := range lines {
		firstAndLast := getFirstAndLastDigits(l)
		calibration, err := strconv.Atoi(firstAndLast)
		if err != nil {
			panic(err)
		}
		sum += calibration
	}
	fmt.Printf("Part 1 answer: %d\n", sum)
}

func getFirstAndLastDigits(s string) string {
	firstDigit := ""
	secondDigit := ""
	for _, r := range s {
		if unicode.IsDigit(r) {
			if firstDigit == "" {
				firstDigit = string(r)
			} else {
				secondDigit = string(r)
			}
		}
	}
	if firstDigit == "" {
		firstDigit = "0"
	}

	if secondDigit == "" {
		secondDigit = firstDigit
	}
	return firstDigit + secondDigit
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	numberMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	for _, l := range lines {
		firstAndLast := getFirstAndLastDigitsReplacingWords(l, numberMap)
		calibration, err := strconv.Atoi(firstAndLast)
		if err != nil {
			panic(err)
		}
		sum += calibration
	}
	fmt.Printf("Part 2 answer: %d\n", sum)

}

func getFirstAndLastDigitsReplacingWords(s string, numMap map[string]string) string {
	digitPositions := map[int]string{}
	for key, val := range numMap {
		idx := strings.Index(s, key)
		if idx != -1 {
			digitPositions[idx] = val

		}
		lastIdx := strings.LastIndex(s, key)
		if lastIdx != -1 {
			digitPositions[lastIdx] = val
		}
	}
	for i, c := range s {
		if unicode.IsDigit(c) {
			digitPositions[i] = string(c)
		}
	}
	var mapKeys []int
	for key := range digitPositions {
		mapKeys = append(mapKeys, key)
	}
	sort.Ints(mapKeys)

	firstIdx := mapKeys[0]
	lastIdx := mapKeys[len(mapKeys)-1]

	firstDigit := digitPositions[firstIdx]
	lastDigit := digitPositions[lastIdx]

	return firstDigit + lastDigit
}
