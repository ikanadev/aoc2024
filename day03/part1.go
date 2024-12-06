package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ikanadev/aoc2024/common"
)

func getMul(s string) int {
	// the string is mul(X,Y) where X and Y are numbers
	numbersStr := s[4 : len(s)-1]
	numberParts := strings.Split(numbersStr, ",")
	x, _ := strconv.Atoi(numberParts[0])
	y, _ := strconv.Atoi(numberParts[1])
	return x * y
}

func Part1() {
	// lines := common.ReadFile("day03/inputSmall")
	lines := common.ReadFile("day03/input")
	total := 0
	for _, line := range lines {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		results := re.FindAll([]byte(line), -1)
		for _, result := range results {
			total += getMul(string(result))
		}
	}
	fmt.Printf("D3P1: %d\n", total)
}

func Part2() {
	// lines := common.ReadFile("day03/inputSmall2")
	lines := common.ReadFile("day03/input")
	// it's just a single line
	line := ""
	for _, l := range lines {
		line += l
	}

	total := 0
	// find the intervals of don't ---> do (which are invalid intervals)
	reDonts := regexp.MustCompile(`don\'t\(\).*?do\(\)`)
	dontIntervals := reDonts.FindAllIndex([]byte(line), -1)

	// create a new string without the don't intervals
	var cleanStr strings.Builder
	lastEnd := 0
	for _, dontInterval := range dontIntervals {
		cleanStr.WriteString(line[lastEnd:dontInterval[0]])
		lastEnd = dontInterval[1]
	}

	// check if there is any remaining don't at the end
	remainingDont := strings.Index(line[lastEnd:], "don't")
	if remainingDont > 0 {
		cleanStr.WriteString(line[lastEnd : lastEnd+remainingDont])
	} else {
		cleanStr.WriteString(line[lastEnd:])
	}
	cleanLine := cleanStr.String()

	// use regex to find all the mul(X,Y)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	results := re.FindAll([]byte(cleanLine), -1)

	for _, result := range results {
		str := string(result)
		total += getMul(str)
	}
	fmt.Printf("D3P2: %d\n", total)
}
