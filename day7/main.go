package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fiveOfAKind = 7
const fourOfAKind = 6
const fullHouse = 5
const threeOfAKind = 4
const twoPair = 3
const onePair = 2
const highCard = 1

func Solve() {
	fmt.Println("Day 7:")
	data, err := os.ReadFile("day7/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

type hand struct {
	cards string
	bet   int
}

func getHands(lines []string) []hand {
	hands := []hand{}
	for _, l := range lines {
		cardsAndBet := strings.Split(l, " ")
		cards := cardsAndBet[0]
		bet, _ := strconv.Atoi(cardsAndBet[1])
		h := hand{cards, bet}
		hands = append(hands, h)
	}
	return hands
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	hands := getHands(lines)
	cardValues := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return compareCards(hands[i], hands[j], cardValues) < 0
	})
	total := 0
	rank := len(hands)
	for i := 0; i < rank; i++ {
		card := hands[i]
		amount := (rank - i) * card.bet
		total += amount
	}
	fmt.Printf("Part 1 answer: %v\n", total)
	// rank hands
	// calculate total winnings
}

func hasFiveOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 5 {
			return true
		}
	}
	return false
}

func hasFourOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 4 {
			return true
		}
	}
	return false
}

func hasFullhouse(cardCounts map[rune]int) bool {
	threeOfAKind := false
	pair := false
	for _, count := range cardCounts {
		if count == 3 {
			threeOfAKind = true
		}
		if count == 2 {
			pair = true
		}
	}
	return threeOfAKind && pair
}

func hasThreeOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 3 {
			return true
		}
	}
	return false
}

func hasTwoPair(cardCounts map[rune]int) bool {
	pair1 := false
	pair2 := false
	for _, count := range cardCounts {
		if count == 2 && pair1 {
			pair2 = true
		} else if count == 2 {
			pair1 = true
		}
	}
	return pair1 && pair2
}

func hasOnePair(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 2 {
			return true
		}
	}
	return false
}

func rankHand(cards string) int {
	counts := cardCounts(cards)
	if hasFiveOfAKind(counts) {
		return fiveOfAKind
	}
	if hasFourOfAKind(counts) {
		return fourOfAKind
	}
	if hasFullhouse(counts) {
		return fullHouse
	}
	if hasThreeOfAKind(counts) {
		return threeOfAKind
	}
	if hasTwoPair(counts) {
		return twoPair
	}
	if hasOnePair(counts) {
		return onePair
	}
	return highCard
}

func compareCardsOfSameRank(c1, c2 string, cardValues map[rune]int) int {
	for i := 0; i < 5; i++ {
		if cardValues[rune(c1[i])] > cardValues[rune(c2[i])] {
			return -1
		} else if cardValues[rune(c1[i])] < cardValues[rune(c2[i])] {
			return 1
		}
	}
	return 0
}

func compareCards(h1, h2 hand, cardValues map[rune]int) int {
	h1Rank := rankHand(h1.cards)
	h2Rank := rankHand(h2.cards)
	if h1Rank > h2Rank {
		return -1
	} else if h1Rank < h2Rank {
		return 1
	}

	return compareCardsOfSameRank(h1.cards, h2.cards, cardValues)
}

func cardCounts(cards string) map[rune]int {
	counts := map[rune]int{}
	for _, card := range cards {
		if _, prs := counts[card]; prs {
			counts[card]++
		} else {
			counts[card] = 1
		}
	}
	return counts
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	hands := getHands(lines)
	cardValues := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return compareCards2(hands[i], hands[j], cardValues) < 0
	})
	total := 0
	rank := len(hands)
	for i := 0; i < rank; i++ {
		card := hands[i]
		amount := (rank - i) * card.bet
		total += amount
	}
	fmt.Printf("Part 2 answer: %v\n", total)
}

func compareCards2(h1, h2 hand, cardValues map[rune]int) int {
	h1Rank := rankHand2(h1.cards)
	h2Rank := rankHand2(h2.cards)
	if h1Rank > h2Rank {
		return -1
	} else if h1Rank < h2Rank {
		return 1
	}

	return compareCardsOfSameRank(h1.cards, h2.cards, cardValues)
}

func hasFiveOfAKind2(cardCounts map[rune]int) bool {
	jokerCount := cardCounts['J']
	for card, count := range cardCounts {
		if card != 'J' && count+jokerCount == 5 {
			return true
		}
	}
	return jokerCount == 5
}

func hasFourOfAKind2(cardCounts map[rune]int) bool {
	jokerCount := cardCounts['J']
	for card, count := range cardCounts {
		if card != 'J' && count+jokerCount == 4 {
			return true
		}
	}
	return jokerCount == 4
}

func hasFullhouse2(cardCounts map[rune]int) bool {
	jackCount := cardCounts['J']
	threeOfAKind := false
	pair := false
	if jackCount == 0 {
		for _, count := range cardCounts {
			if count == 3 {
				threeOfAKind = true
			} else if count == 2 {
				pair = true
			}
		}
	} else {
		var pair1 = '0'
		for card, count := range cardCounts {
			if pair1 == '0' && count == 2 {
				pair1 = card
				threeOfAKind = true
			} else if card != pair1 && count == 2 {
				pair = true
			}
		}
	}

	return threeOfAKind && pair
}

func hasThreeOfAKind2(cardCounts map[rune]int) bool {
	jokerCount := cardCounts['J']
	for card, count := range cardCounts {
		if card != 'J' && count+jokerCount == 3 {
			return true
		}
	}
	return jokerCount == 3
}

func hasOnePair2(cardCounts map[rune]int) bool {
	jokerCount := cardCounts['J']
	for card, count := range cardCounts {
		if card != 'J' && count+jokerCount == 2 {
			return true
		}
	}
	return jokerCount == 'J'
}

func rankHand2(cards string) int {
	counts := cardCounts(cards)
	if hasFiveOfAKind2(counts) {
		return fiveOfAKind
	}
	if hasFourOfAKind2(counts) {
		return fourOfAKind
	}
	if hasFullhouse2(counts) {
		return fullHouse
	}
	if hasThreeOfAKind2(counts) {
		return threeOfAKind
	}
	if hasTwoPair(counts) {
		return twoPair
	}
	if hasOnePair2(counts) {
		return onePair
	}
	return highCard
}
