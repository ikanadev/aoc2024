package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ikanadev/aoc2024/common"
)

func sortAsc(a, b int) int {
	return a - b
}

func Part1() {
	// lines := common.ReadFile("day01/inputSmall")
	lines := common.ReadFile("day01/input")
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))
	for i := range lines {
		numbersStr := strings.Split(lines[i], "   ")
		leftList[i], _ = strconv.Atoi(numbersStr[0])
		rightList[i], _ = strconv.Atoi(numbersStr[1])
	}
	slices.SortFunc(leftList, sortAsc)
	slices.SortFunc(rightList, sortAsc)
	total := 0
	for i := range leftList {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}
	fmt.Printf("D1P1: %d\n", total)
}
