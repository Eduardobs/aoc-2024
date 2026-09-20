package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var instructionPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func FirstProblem() (int64, error) {
	return Solution1()
}

func SecondProblem() (int64, error) {
	return Solution2()
}

func Solution1() (int64, error) {
	content, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 3 input: %w", err)
	}
	return solve(string(content), false), nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 3 input: %w", err)
	}
	return solve(string(content), true), nil
}

func solve(input string, conditionals bool) int64 {
	res := instructionPattern.FindAllStringSubmatch(input, -1)
	var sum int64
	enabled := true
	for _, l := range res {
		switch l[0] {
		case "do()":
			if conditionals {
				enabled = true
			}
		case "don't()":
			if conditionals {
				enabled = false
			}
		default:
			if enabled {
				n1, _ := strconv.Atoi(l[1])
				n2, _ := strconv.Atoi(l[2])
				sum += int64(n1 * n2)
			}
		}
	}

	return sum
}
