package day5

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve() {
	fmt.Println("Day 5:")
	data, err := os.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	part1(input)
	part2(input)
}

type mapInfo struct {
	destinationStart int
	sourceStart      int
	length           int
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines)

	maps := createMaps(lines)
	locations := []int{}
	for _, s := range seeds {
		loc := seedLocation(s, maps)
		locations = append(locations, loc)
	}
	minLocation := min(locations)
	fmt.Printf("Part 1 answer: %v\n", minLocation)

}

func parseSeeds(lines []string) []int {
	seedLine := lines[0]
	labelAndSeeds := strings.Split(seedLine, ":")
	seedsString := labelAndSeeds[1]
	seedStrings := strings.Fields(seedsString)
	seeds := []int{}
	for _, s := range seedStrings {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}
	return seeds
}

func createMaps(lines []string) map[int][]mapInfo {
	maps := map[int][]mapInfo{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
	}

	mapNum := 0
	for _, l := range lines[2:] {
		if len(strings.TrimSpace(l)) == 0 {
			mapNum++
		} else if unicode.IsDigit(rune(l[0])) {
			mapNumStrings := strings.Fields(l)
			mapNums := []int{}
			for _, r := range mapNumStrings {
				n, _ := strconv.Atoi(string(r))
				mapNums = append(mapNums, n)
			}
			info := mapInfo{destinationStart: mapNums[0], sourceStart: mapNums[1], length: mapNums[2]}
			maps[mapNum] = append(maps[mapNum], info)
		}
	}
	return maps
}

func seedLocation(seed int, maps map[int][]mapInfo) int {
	value := seed
	mapNum := 0
	for mapNum < len(maps) {
		value = convert(value, maps[mapNum])
		mapNum++
	}
	return value
}

func convert(value int, conversions []mapInfo) int {
	for _, v := range conversions {
		if value >= v.sourceStart && value <= v.sourceStart+v.length {
			diff := value - v.sourceStart
			return v.destinationStart + diff
		}
	}
	return value
}

func min(nums []int) int {
	m := math.MaxInt
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

type seedRange struct {
	start int
	end   int
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	seedInfo := parseSeeds(lines)
	seedRanges := []seedRange{}
	for i := 0; i < len(seedInfo)-1; i += 2 {
		start, length := seedInfo[i], seedInfo[i+1]
		seedRange := seedRange{start: start, end: start + length - 1}
		seedRanges = append(seedRanges, seedRange)
	}
	maps := createMaps(lines)
	answer := 999999999999999
	for _, r := range seedRanges {
		loc := seedRangeLocation(r, maps)
		if loc < answer {
			answer = loc
		}
	}
	fmt.Printf("Part 2 answer: %v\n", answer)

}

func mappedRanges(seeds []seedRange, maps []mapInfo) []seedRange {
	rangesToMap := seeds
	result := []seedRange{}
	for _, m := range maps {
		next := []seedRange{}
		for _, r := range rangesToMap {
			before, mappedOverlap, after := mapRange(r, m)
			if before.start != -1 {
				next = append(next, before)
			}
			if mappedOverlap.start != -1 {
				result = append(result, mappedOverlap)
			}
			if after.start != -1 {
				next = append(next, after)
			}
		}
		rangesToMap = next
	}
	result = append(result, rangesToMap...)
	return result
}

func mapRange(r seedRange, m mapInfo) (seedRange, seedRange, seedRange) {
	mapSourceEnd := m.sourceStart + m.length - 1
	// before case
	before := seedRange{-1, -1}
	if r.start < m.sourceStart {
		before.start = r.start
		before.end = minNum(r.end, m.sourceStart-1)
	}
	overlap := seedRange{-1, -1}
	// overlap case
	if m.sourceStart < r.end && mapSourceEnd > r.start {
		overlap.start = maxNum(m.sourceStart, r.start)
		overlap.end = minNum(mapSourceEnd, r.end)
		startDiff := overlap.start - m.sourceStart
		endDiff := overlap.end - m.sourceStart
		overlap.start = m.destinationStart + startDiff
		overlap.end = m.destinationStart + endDiff
	}
	// after case
	after := seedRange{-1, -1}
	if r.end > mapSourceEnd {
		after.start = maxNum(r.start, mapSourceEnd+1)
		after.end = r.end
	}
	return before, overlap, after
}

func minNum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxNum(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func seedRangeLocation(seeds seedRange, maps map[int][]mapInfo) int {
	seedRanges := []seedRange{seeds}
	mapNum := 0
	for mapNum < len(maps) {
		// convert range
		currentMap := maps[mapNum]
		seedRanges = mappedRanges(seedRanges, currentMap)
		mapNum++
	}
	result := lowestLocation(seedRanges)
	return result
}

func lowestLocation(seedRanges []seedRange) int {
	lowest := 99999999999999999
	for _, r := range seedRanges {
		if r.start < lowest {
			lowest = r.start
		}
	}
	return lowest
}
