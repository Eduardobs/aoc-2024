package day03

import (
	"errors"
	"os"
	"regexp"
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
	content, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	text := ""
	for i := 0; i < len(lines); i++ {
		text += lines[i]
	}

	rg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	res := rg.FindAllStringSubmatch(text, -1)

	sum := 0
	for _, l := range res {
		n1, _ := strconv.Atoi(l[1])
		n2, _ := strconv.Atoi(l[2])
		sum += n1 * n2
	}

	return int64(sum), nil
}

func Solution2() (int64, error) {
	content, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	text := ""
	for i := 0; i < len(lines); i++ {
		text += lines[i]
	}

	rg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

	res := rg.FindAllStringSubmatch(text, -1)

	sum := 0
	enabled := true
	for _, l := range res {
		switch {
		case strings.HasPrefix(l[0], "mul"):
			if enabled {
				n1, _ := strconv.Atoi(l[1])
				n2, _ := strconv.Atoi(l[2])
				sum += n1 * n2
			}

		case l[0] == "do()":
			enabled = true

		case l[0] == "don't()":
			enabled = false
		}
	}

	return int64(sum), nil
}
