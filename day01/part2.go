package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ikanadev/aoc2024/common"
)

func Part2() {
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
	total := 0
	scoresMap := make(map[int]int)
	for i := range leftList {
		numberToCount := leftList[i]
		score, ok := scoresMap[numberToCount]
		if ok {
			total += score
			continue
		}
		timesFound := 0
		for j := range rightList {
			if numberToCount == rightList[j] {
				timesFound++
			}
		}
		scoresMap[numberToCount] = numberToCount * timesFound
		total += numberToCount * timesFound
	}
	fmt.Printf("Part2: %d\n", total)
}
