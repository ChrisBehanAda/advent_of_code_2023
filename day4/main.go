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
	part2(input)
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
	fmt.Printf("Part 1 answer: %v\n", totalPoints)
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
	winningNums := winningNums(card)
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

func winningNums(card scratchCard) int {
	winningNums := 0
	for _, p := range card.playerNumbers {
		for _, w := range card.winningNumbers {
			if p == w {
				winningNums++
				break
			}
		}
	}
	return winningNums
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	cardMap := map[int][]scratchCard{}
	for i, l := range lines {
		card := parseScratchCard(l)
		cardMap[i+1] = []scratchCard{card}
	}

	for i := 1; i <= len(cardMap); i++ {
		cards := cardMap[i]
		firstCard := cards[0]
		matchCount := winningNums(firstCard)
		copies := len(cards)
		for k := 0; k < copies; k++ {
			for j := i + 1; j <= i+matchCount; j++ {
				cardToCopy := cardMap[j][0]
				winningNumsCopy := make([]int, len(cardToCopy.winningNumbers))
				copy(winningNumsCopy, cardToCopy.winningNumbers)
				playerNumsCopy := make([]int, len(cardToCopy.playerNumbers))
				copy(playerNumsCopy, cardToCopy.playerNumbers)
				cardCopy := scratchCard{winningNumbers: winningNumsCopy, playerNumbers: playerNumsCopy}
				cardMap[j] = append(cardMap[j], cardCopy)
			}
		}
		// produce a single copy for the next matchCount number of cards
	}
	totalCards := 0
	for _, cards := range cardMap {
		totalCards += len(cards)
	}
	fmt.Printf("Part 2 answer: %v\n", totalCards)
}
