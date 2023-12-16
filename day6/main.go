package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve() {
	fmt.Println("Day 6:")
	data, err := os.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

type Record struct {
	time     int
	distance int
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	timeLine, distanceLine := lines[0], lines[1]
	timeNumberString := strings.Split(timeLine, ":")[1]
	distanceNumberString := strings.Split(distanceLine, ":")[1]
	timeStrings := strings.Fields(timeNumberString)
	distanceStrings := strings.Fields(distanceNumberString)
	records := []Record{}

	for i := 0; i < len(timeStrings); i++ {
		timeString := timeStrings[i]
		distanceString := distanceStrings[i]
		time, _ := strconv.Atoi(timeString)
		distance, _ := strconv.Atoi(distanceString)
		r := Record{time: time, distance: distance}
		records = append(records, r)
	}

	wins := waysToWin(records)
	ans := 1
	for _, w := range wins {
		ans *= w
	}
	fmt.Printf("Part 1 answer: %v\n", ans)
}

func winCount(r Record) int {
	count := 0
	for t := 0; t <= r.time; t++ {
		distance := distanceTraveled(t, r.time)
		if distance > r.distance {
			count++
		}
	}
	return count
}

func waysToWin(records []Record) []int {
	ways := []int{}
	for _, r := range records {
		count := winCount(r)
		ways = append(ways, count)
	}

	return ways
}

func distanceTraveled(secondsHeld int, total int) int {
	millimetersPerSecond := secondsHeld
	timeToTravel := total - secondsHeld
	distance := millimetersPerSecond * timeToTravel
	return distance
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	timeLine, distanceLine := lines[0], lines[1]
	timeNumberString := strings.Split(timeLine, ":")[1]
	distanceNumberString := strings.Split(distanceLine, ":")[1]
	timeNumberString = removeWhitespace(timeNumberString)
	distanceNumberString = removeWhitespace(distanceNumberString)

	time, _ := strconv.Atoi(timeNumberString)
	distance, _ := strconv.Atoi(distanceNumberString)
	record := Record{time, distance}
	ans := winCount(record)
	fmt.Printf("Part 2 ans: %v\n", ans)

}

func removeWhitespace(s string) string {
	resultRunes := []rune{}
	for _, r := range s {
		if !unicode.IsSpace(r) {
			resultRunes = append(resultRunes, r)
		}
	}
	return string(resultRunes)
}
