package day04

import (
	"errors"
	"os"
	"strings"
)

func FirstProblem() (int64, error) {
	return Solution1()
}

func SecondProblem() (int64, error) {
	return Solution1()
}

func Solution1() (int64, error) {
	content, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	word := "XMAS"
	rows := len(lines)
	cols := len(lines[0])
	wordLen := len(word)
	count := 0

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

	return int64(count), nil
}
