package day04

import "github.com/ikanadev/aoc2024/common"

type Matrix struct {
	cols int
	rows int
	data []string
}

func (m Matrix) isValidCell(x, y int) bool {
	return x >= 0 && x < m.cols && y >= 0 && y < m.rows
}

func (m Matrix) checkBottomRight(x, y int, word string) bool {
	if !m.isValidCell(x+len(word)-1, y+len(word)-1) {
		return false
	}
	for k, i := 0, x; i < x+len(word); k, i = k+1, i+1 {
		j := y + k
		if m.data[i][j] != word[k] {
			return false
		}
	}
	return true
}

func (m Matrix) checkRight(x, y int, word string) bool {
	if !m.isValidCell(x, y+len(word)-1) {
		return false
	}
	return m.data[x][y:y+len(word)] == word
}

func (m Matrix) checkLeft(x, y int, word string) bool {
	if !m.isValidCell(x, y-len(word)+1) {
		return false
	}
	for k, j := 0, y; j > y-len(word); k, j = k+1, j-1 {
		if m.data[x][j] != word[k] {
			return false
		}
	}
	return true
}

func (m Matrix) checkBottom(x, y int, word string) bool {
	if !m.isValidCell(x+len(word)-1, y) {
		return false
	}
	for k, i := 0, x; i < x+len(word); k, i = k+1, i+1 {
		if m.data[i][y] != word[k] {
			return false
		}
	}
	return true
}

// UNCHECKED

func (m Matrix) checkTop(x, y int, word string) bool {
	if !m.isValidCell(x-len(word)+1, y) {
		return false
	}
	for k, i := 0, x; i > x-len(word); k, i = k+1, i-1 {
		if m.data[i][y] != word[k] {
			return false
		}
	}
	return true
}

func (m Matrix) checkTopLeft(x, y int, word string) bool {
	if !m.isValidCell(x-len(word)+1, y-len(word)+1) {
		return false
	}
	for k, i := 0, x; i > x-len(word); k, i = k+1, i-1 {
		for j := y; j > y-len(word); j = j - 1 {
			if m.data[i][j] != word[k] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) checkTopRight(x, y int, word string) bool {
	if !m.isValidCell(x+len(word)-1, y-len(word)+1) {
		return false
	}
	for k, i := 0, x; i < x+len(word); k, i = k+1, i+1 {
		for j := y; j > y-len(word); j = j - 1 {
			if m.data[i][j] != word[k] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) checkBottomLeft(x, y int, word string) bool {
	if !m.isValidCell(x-len(word)+1, y+len(word)-1) {
		return false
	}
	for k, i := 0, x; i > x-len(word); k, i = k+1, i-1 {
		for j := y; j < y+len(word); j = j + 1 {
			if m.data[i][j] != word[k] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) countAnyDirection(i, j int, word string) int {
	counter := 0
	if m.checkRight(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkLeft(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkTop(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkBottom(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkTopLeft(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkTopRight(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkBottomLeft(i, j, word) {
		println("found:", i, j)
		counter++
	}
	if m.checkBottomRight(i, j, word) {
		println("found:", i, j)
		counter++
	}
	return counter
}

func (m Matrix) checkXmas() {
	word := "XMAS"
	counter := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			counter += m.countAnyDirection(i, j, word)
		}
	}
	println(counter)
}

func Part1() {
	lines := common.ReadFile("day04/inputSmall")
	matrix := Matrix{
		cols: len(lines[0]),
		rows: len(lines),
		data: lines,
	}
	print(matrix.checkBottomLeft(3, 9, "XMAS"))
}
