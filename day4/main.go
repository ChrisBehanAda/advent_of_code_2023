package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day 4:")
	data, err := os.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
}

type scratchCard struct {
	winningNumbers []int
	playerNumbers  []int
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	totalPoints := 0
	for _, l := range lines {
		scratchCard := parseScratchCard(l)
		points := score(scratchCard)
		totalPoints += points
	}
	fmt.Printf("Part 1 answer: %v", totalPoints)
}

func parseScratchCard(line string) scratchCard {
	cardInfoAndNums := strings.Split(line, ":")
	nums := cardInfoAndNums[1]
	winningAndPlayerNums := strings.Split(nums, "|")
	winningNumsString := strings.TrimSpace(winningAndPlayerNums[0])
	playerNumsString := strings.TrimSpace(winningAndPlayerNums[1])
	winningNumStrings := strings.Fields(winningNumsString)
	playerNumStrings := strings.Fields(playerNumsString)

	winningNums := []int{}
	for _, s := range winningNumStrings {
		num, _ := strconv.Atoi(s)
		winningNums = append(winningNums, num)
	}

	playerNums := []int{}
	for _, s := range playerNumStrings {
		num, _ := strconv.Atoi(s)
		playerNums = append(playerNums, num)
	}

	return scratchCard{winningNumbers: winningNums, playerNumbers: playerNums}
}

func score(card scratchCard) int {
	winningNums := 0
	for _, p := range card.playerNumbers {
		for _, w := range card.winningNumbers {
			if p == w {
				winningNums++
				break
			}
		}
	}
	score := 0
	for i := 0; i < winningNums; i++ {
		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score
}
