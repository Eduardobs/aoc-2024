package day02

import (
	"fmt"
	"os"
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
	content, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 2 input: %w", err)
	}
	return solve(string(content), false)
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		return 0, fmt.Errorf("read day 2 input: %w", err)
	}
	return solve(string(content), true)
}

func solve(input string, dampener bool) (int64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, nil
	}
	lines := strings.Split(input, "\n")

	var count int64

	for i, line := range lines {
		list, err := toInt(strings.Fields(line))
		if err != nil {
			return 0, fmt.Errorf("day 2 line %d: %w", i+1, err)
		}

		if safe(list) || dampener && safeByRemoving(list) {
			count++
		}
	}
	return count, nil
}

func toInt(xs []string) ([]int, error) {
	l := make([]int, len(xs))
	for i, x := range xs {
		v, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		l[i] = v
	}
	return l, nil
}

func safe(xs []int) bool {
	if len(xs) < 2 {
		return true
	}

	asc := xs[1]-xs[0] > 0

	for i := 1; i < len(xs); i++ {
		incr := xs[i] - xs[i-1]

		if incr > 0 != asc {
			return false
		}

		if incr < 0 {
			incr = -incr
		}

		if incr < 1 || incr > 3 {
			return false
		}
	}
	return true
}

func safeByRemoving(list []int) bool {
	for i := 0; i < len(list); i++ {
		newXs := append([]int{}, list[:i]...)
		newXs = append(newXs, list[i+1:]...)

		if safe(newXs) {
			return true
		}
	}
	return false
}
