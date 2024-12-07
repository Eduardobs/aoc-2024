package day01

import (
	"errors"
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
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	var a, b []int
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], "   ")

		n1, _ := strconv.Atoi(line[0])
		n2, _ := strconv.Atoi(strings.TrimSuffix(line[1], "\r"))

		a = append(a, n1)
		b = append(b, n2)
	}

	sort.Ints(a)
	sort.Ints(b)

	diff := 0
	for i := 0; i < len(a); i++ {
		diff += abs(a[i] - b[i])
	}

	return int64(diff), nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	var a []int
	var b = map[int]int{}

	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], "   ")

		n1, _ := strconv.Atoi(line[0])
		n2, _ := strconv.Atoi(strings.TrimSuffix(line[1], "\r"))

		a = append(a, n1)
		b[n2]++
	}

	simScore := 0
	for _, n := range a {
		simScore += n * b[n]
	}

	return int64(simScore), nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
