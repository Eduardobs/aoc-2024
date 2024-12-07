package day02

import (
	"errors"
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
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	count := 0

	for i := 0; i < len(lines); i++ {
		list, _ := toInt(strings.Fields(lines[i]))

		if safe(list) {
			count++
		}
	}
	return int64(count), nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	count := 0

	for i := 0; i < len(lines); i++ {
		list, _ := toInt(strings.Fields(lines[i]))

		if safe(list) || safeByRemoving(list) {
			count++
		}
	}
	return int64(count), nil
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
