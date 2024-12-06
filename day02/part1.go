package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ikanadev/aoc2024/common"
)

func isReportSafe(report []int) bool {
	increasing := true
	if report[0] > report[1] {
		increasing = false
	}

	last := report[0]
	for i := 1; i < len(report); i++ {
		if increasing {
			if last >= report[i] || report[i]-last > 3 {
				return false
			}
		} else {
			if last <= report[i] || last-report[i] > 3 {
				return false
			}
		}
		last = report[i]
	}
	return true
}

// call only when a report is unsafe
func checkSafeRemovingOne(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}

func Part1() {
	// lines := common.ReadFile("day02/inputSmall")
	lines := common.ReadFile("day02/input")
	totalSafe := 0
	for i := range lines {
		numbersStr := strings.Split(lines[i], " ")
		numbers := make([]int, len(numbersStr))
		for i := range numbersStr {
			numbers[i], _ = strconv.Atoi(numbersStr[i])
		}
		if isReportSafe(numbers) {
			totalSafe++
		}
	}
	fmt.Printf("D2P1: %d\n", totalSafe)
}

func Part2() {
	// lines := common.ReadFile("day02/inputSmall")
	lines := common.ReadFile("day02/input")
	totalSafe := 0
	for i := range lines {
		numbersStr := strings.Split(lines[i], " ")
		numbers := make([]int, len(numbersStr))
		for i := range numbersStr {
			numbers[i], _ = strconv.Atoi(numbersStr[i])
		}
		if isReportSafe(numbers) {
			totalSafe++
		} else {
			if checkSafeRemovingOne(numbers) {
				totalSafe++
			}
		}
	}
	fmt.Printf("D2P2: %d\n", totalSafe)
}
