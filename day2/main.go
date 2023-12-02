package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type result struct {
	red   int
	green int
	blue  int
}

// Solve is the entry point for running the solutions.
func Solve() {
	fmt.Println("Day 2:")
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	possibleGameIDSum := 0
	for _, l := range lines {
		id, game := getGameResults(l)
		if isPossibleGame(game) {
			possibleGameIDSum += id
		}
	}
	fmt.Printf("Part 1 answer: %v\n", possibleGameIDSum)
}

func getGameResults(line string) (int, []result) {
	gameAndRounds := strings.SplitN(line, ":", 2)
	gameString := gameAndRounds[0]
	roundsString := gameAndRounds[1]
	gameIDString := strings.Split(gameString, " ")[1]
	gameID, _ := strconv.Atoi(gameIDString)

	results := []result{}
	rounds := strings.Split(roundsString, ";")
	redRegex := regexp.MustCompile(`\s*(\d+) red`)
	greenRegex := regexp.MustCompile(`\s*(\d+) green`)
	blueRegex := regexp.MustCompile(`\s*(\d+) blue`)
	for _, r := range rounds {
		cubeColors := result{}
		redMatch := redRegex.FindStringSubmatch(r)
		if redMatch != nil {
			redCubes, _ := strconv.Atoi(redMatch[1])
			cubeColors.red = redCubes
		}
		blueMatch := blueRegex.FindStringSubmatch(r)
		if blueMatch != nil {
			blueCubes, _ := strconv.Atoi(blueMatch[1])
			cubeColors.blue = blueCubes
		}
		greenMatch := greenRegex.FindStringSubmatch(r)
		if greenMatch != nil {
			greenCubes, _ := strconv.Atoi(greenMatch[1])
			cubeColors.green = greenCubes
		}
		results = append(results, cubeColors)
	}
	return gameID, results
}

func isPossibleGame(game []result) bool {
	maxRed, maxBlue, maxGreen := 12, 14, 13
	for _, r := range game {
		if r.red > maxRed || r.blue > maxBlue || r.green > maxGreen {
			return false
		}
	}
	return true
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	powerSum := 0
	for _, l := range lines {
		_, game := getGameResults(l)
		power := powerOfCubes(game)
		powerSum += power
	}
	fmt.Printf("Part 2 answer: %v\n", powerSum)
}

func powerOfCubes(game []result) int {
	maxRed, maxBlue, maxGreen := 1, 1, 1 // initialize to 1 so multiplication works
	for _, g := range game {
		if g.red > maxRed {
			maxRed = g.red
		}
		if g.blue > maxBlue {
			maxBlue = g.blue
		}
		if g.green > maxGreen {
			maxGreen = g.green
		}
	}
	return maxRed * maxBlue * maxGreen
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
