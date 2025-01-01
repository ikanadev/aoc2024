package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ikanadev/aoc2024/common"
)

type rules map[int][]int

func parseRules(lines []string) rules {
	data := make(rules)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		if _, ok := data[from]; !ok {
			data[from] = make([]int, 0)
		}
		data[from] = append(data[from], to)
	}
	return data
}

func Part1() {
	// lines := common.ReadFile("day05/inputSmall")
	lines := common.ReadFile("day05/input")
	blankLineIndex := 0
	for i, line := range lines {
		if line == "" {
			blankLineIndex = i
		}
	}
	rulesLines := lines[:blankLineIndex]
	updatesLines := lines[blankLineIndex+1:]

	rules := parseRules(rulesLines)
	updates := make([][]int, 0)
	for _, line := range updatesLines {
		parts := strings.Split(line, ",")
		numbers := make([]int, len(parts))
		for i := range parts {
			number, _ := strconv.Atoi(parts[i])
			numbers[i] = number
		}
		updates = append(updates, numbers)
	}

	total := 0
	for _, updateNumbers := range updates {
		isValid := true
		for i, number := range updateNumbers {
			if i == len(updateNumbers)-1 {
				continue
			}
			if _, ok := rules[number]; !ok {
				isValid = false
				break
			}
			allowedNumbers := rules[number]

			for j := i + 1; j < len(updateNumbers); j++ {
				if !slices.Contains(allowedNumbers, updateNumbers[j]) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			total += updateNumbers[len(updateNumbers)/2]
		}
	}

	fmt.Printf("D04P1: %d\n", total)
}
