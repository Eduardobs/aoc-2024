package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstProblem() (int64, error) {
	return Solution1()
}

func SecondProblem() (int64, error) {
	return Solution2()
}

func Solution1() (int64, error) {
	content, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 1 input: %w", err)
	}

	a, b, err := parseLists(string(content))
	if err != nil {
		return 0, err
	}
	return totalDistance(a, b), nil
}

func totalDistance(a, b []int) int64 {
	sort.Ints(a)
	sort.Ints(b)

	var diff int64
	for i := range a {
		diff += int64(abs(a[i] - b[i]))
	}

	return diff
}

func parseLists(input string) ([]int, []int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var a, b []int
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("day 1 line %d: expected two numbers", i+1)
		}

		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("day 1 line %d: %w", i+1, err)
		}
		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("day 1 line %d: %w", i+1, err)
		}

		a = append(a, n1)
		b = append(b, n2)
	}

	return a, b, nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 1 input: %w", err)
	}

	a, right, err := parseLists(string(content))
	if err != nil {
		return 0, err
	}
	return similarityScore(a, right), nil
}

func similarityScore(left, right []int) int64 {
	counts := make(map[int]int, len(right))
	for _, n := range right {
		counts[n]++
	}

	var simScore int64
	for _, n := range left {
		simScore += int64(n * counts[n])
	}

	return simScore
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
