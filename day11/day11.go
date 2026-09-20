package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key struct {
	n     int64
	steps int
}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day11/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 11 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	fs := strings.Fields(input)
	nums := make([]int64, len(fs))
	for i, s := range fs {
		n, e := strconv.ParseInt(s, 10, 64)
		if e != nil {
			return 0, 0, e
		}
		nums[i] = n
	}
	memo := map[key]int64{}
	var count func(int64, int) int64
	count = func(n int64, steps int) int64 {
		if steps == 0 {
			return 1
		}
		k := key{n, steps}
		if v, ok := memo[k]; ok {
			return v
		}
		var v int64
		if n == 0 {
			v = count(1, steps-1)
		} else {
			s := strconv.FormatInt(n, 10)
			if len(s)%2 == 0 {
				a, _ := strconv.ParseInt(s[:len(s)/2], 10, 64)
				b, _ := strconv.ParseInt(s[len(s)/2:], 10, 64)
				v = count(a, steps-1) + count(b, steps-1)
			} else {
				v = count(n*2024, steps-1)
			}
		}
		memo[k] = v
		return v
	}
	var a, b int64
	for _, n := range nums {
		a += count(n, 25)
		b += count(n, 75)
	}
	return a, b, nil
}
