package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coordinate struct {
	row int
	col int
}

func Solve() {
	fmt.Println("Day 3:")
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	schematic := parseSchematic(lines)
	partNums := partNumbers(schematic)
	partNumSums := 0
	for _, n := range partNums {
		partNumSums += n
	}
	fmt.Printf("Part 1 answer: %v\n", partNumSums)
}

func parseSchematic(lines []string) [][]rune {
	schematic := [][]rune{}
	for _, l := range lines {
		symbols := []rune(l)
		schematic = append(schematic, symbols)
	}
	return schematic
}

func partNumbers(schematic [][]rune) []int {
	// Iterate through each symbol
	// if symbol is a number, create a new number string, and a part number boolean
	// check all adjacent points for a symbol (non digit or .)
	// if one is found, set part number boolean to true
	// continue iterating, adding digit strings to the right to the number
	partNums := []string{}
	for row := 0; row < len(schematic); row++ {
		isNum := false
		num := ""
		part := false
		for col := 0; col < len(schematic[0]); col++ {
			if !isNum {
				if unicode.IsDigit(schematic[row][col]) {
					isNum = true
					num += string(schematic[row][col])
					if isPart(schematic, row, col) {
						part = true
					}
				}
			} else {
				// previous string was a number but this one isn't
				if !unicode.IsDigit(schematic[row][col]) {
					isNum = false
					if part {
						partNums = append(partNums, num)
					}
					num = ""
					part = false
				} else {
					// previous string was a number and so is this one
					num += string(schematic[row][col])
					if isPart(schematic, row, col) {
						part = true
					}
				}
			}
		}
		if isNum && part {
			partNums = append(partNums, num)
		}
	}
	partNumInts := []int{}
	for _, n := range partNums {
		i, _ := strconv.Atoi(string(n))
		partNumInts = append(partNumInts, i)
	}
	return partNumInts
}

func isPart(schematic [][]rune, row, col int) bool {
	leftAboveDiagonal := coordinate{row - 1, col - 1}
	above := coordinate{row - 1, col}
	rightAboveDiagonal := coordinate{row - 1, col + 1}
	left := coordinate{row, col - 1}
	right := coordinate{row, col + 1}
	leftBelowDiagonal := coordinate{row + 1, col - 1}
	below := coordinate{row + 1, col}
	rightBelowDiagonal := coordinate{row + 1, col + 1}
	adjacentPoints := []coordinate{leftAboveDiagonal, above, rightAboveDiagonal, left, right, leftBelowDiagonal, below, rightBelowDiagonal}
	for _, p := range adjacentPoints {
		if p.row >= 0 && p.row < len(schematic) && p.col >= 0 && p.col < len(schematic[0]) {
			if schematic[p.row][p.col] != '.' && !unicode.IsDigit(schematic[p.row][p.col]) {
				return true
			}
		}
	}
	return false
}

// gear has 2 part numbers
type gear struct {
	p1 coordinate
	p2 coordinate
}

type partNum struct {
	n     int
	start coordinate
	end   coordinate
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	schematic := parseSchematic(lines)
	parts := partNumInfo(schematic)
	gearRatioSum := findGears(schematic, parts)
	fmt.Printf("Part 2 answer: %v\n", gearRatioSum)
}

func partNumInfo(schematic [][]rune) []partNum {
	partNums := []partNum{}
	for row := 0; row < len(schematic); row++ {
		isNum := false
		num := ""
		part := false
		start := coordinate{}
		end := coordinate{}
		for col := 0; col < len(schematic[0]); col++ {
			if !isNum {
				if unicode.IsDigit(schematic[row][col]) {
					isNum = true
					num += string(schematic[row][col])
					start.row = row
					start.col = col
					end.row = row
					end.col = col
					if isPart(schematic, row, col) {
						part = true
					}
				}
			} else {
				// previous string was a number but this one isn't
				if !unicode.IsDigit(schematic[row][col]) {
					isNum = false
					if part {
						n, _ := strconv.Atoi(num)
						partNum := partNum{n, start, end}
						partNums = append(partNums, partNum)
					}
					num = ""
					part = false
				} else {
					// previous string was a number and so is this one
					end.row = row
					end.col = col
					num += string(schematic[row][col])
					if isPart(schematic, row, col) {
						part = true
					}
				}
			}
		}
		if isNum && part {
			n, _ := strconv.Atoi(num)
			partNum := partNum{n, start, end}
			partNums = append(partNums, partNum)
		}
	}
	return partNums
}

func findGears(schematic [][]rune, parts []partNum) int {
	gearRatioSum := 0
	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[0]); col++ {
			if schematic[row][col] == '*' {
				isGear, ratio := isGear2(parts, schematic, row, col)
				if isGear {
					gearRatioSum += ratio
				}
			}
		}
	}
	return gearRatioSum
}

func isGear2(parts []partNum, schematic [][]rune, row, col int) (bool, int) {
	leftAboveDiagonal := coordinate{row - 1, col - 1}
	above := coordinate{row - 1, col}
	rightAboveDiagonal := coordinate{row - 1, col + 1}
	left := coordinate{row, col - 1}
	right := coordinate{row, col + 1}
	leftBelowDiagonal := coordinate{row + 1, col - 1}
	below := coordinate{row + 1, col}
	rightBelowDiagonal := coordinate{row + 1, col + 1}
	adjacentPoints := []coordinate{leftAboveDiagonal, above, rightAboveDiagonal, left, right, leftBelowDiagonal, below, rightBelowDiagonal}
	adjacentParts := map[int]bool{}
	for _, adjPoint := range adjacentPoints {
		for idx, p := range parts {
			if p.start.row == adjPoint.row && p.start.col <= adjPoint.col && p.end.col >= adjPoint.col {
				_, prs := adjacentParts[idx]
				if !prs {
					adjacentParts[idx] = true
				}
			}
		}
	}
	if len(adjacentParts) == 2 {
		p1Idx := -1
		p2Idx := -1
		for k := range adjacentParts {
			if p1Idx == -1 {
				p1Idx = k
			} else {
				p2Idx = k
			}
		}
		p1Num := parts[p1Idx].n
		p2Num := parts[p2Idx].n
		return true, p1Num * p2Num
	}
	return false, -1
}
