package day04

import (
	"fmt"
	"os"
	"strings"
)

func FirstProblem() (int64, error) {
	return Solution1()
}

func SecondProblem() (int64, error) {
	return Solution2()
}

func Solution1() (int64, error) {
	content, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 4 input: %w", err)
	}
	lines, err := parseGrid(string(content))
	if err != nil {
		return 0, err
	}
	return countXMAS(lines), nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 4 input: %w", err)
	}
	lines, err := parseGrid(string(content))
	if err != nil {
		return 0, err
	}
	return countXMASCrosses(lines), nil
}

func parseGrid(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("day 4 input is empty")
	}
	lines := strings.Split(input, "\n")
	width := len(strings.TrimSuffix(lines[0], "\r"))
	for i := range lines {
		lines[i] = strings.TrimSuffix(lines[i], "\r")
		if len(lines[i]) != width {
			return nil, fmt.Errorf("day 4 line %d: expected width %d, got %d", i+1, width, len(lines[i]))
		}
	}
	return lines, nil
}

func countXMAS(lines []string) int64 {
	word := "XMAS"
	rows := len(lines)
	cols := len(lines[0])
	wordLen := len(word)
	var count int64

	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{1, 1},   // Down-Right diagonal
		{1, -1},  // Down-Left diagonal
		{0, -1},  // Left
		{-1, 0},  // Up
		{-1, -1}, // Up-Left diagonal
		{-1, 1},  // Up-Right diagonal
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				match := true
				for i := 0; i < wordLen; i++ {
					nr := r + dir[0]*i
					nc := c + dir[1]*i
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || lines[nr][nc] != word[i] {
						match = false
						break
					}
				}
				if match {
					count++
				}
			}
		}
	}

	return count
}

func countXMASCrosses(lines []string) int64 {
	var count int64
	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			if lines[row][col] != 'A' {
				continue
			}
			leftDiagonal := isMASPair(lines[row-1][col-1], lines[row+1][col+1])
			rightDiagonal := isMASPair(lines[row-1][col+1], lines[row+1][col-1])
			if leftDiagonal && rightDiagonal {
				count++
			}
		}
	}
	return count
}

func isMASPair(a, b byte) bool {
	return a == 'M' && b == 'S' || a == 'S' && b == 'M'
}
