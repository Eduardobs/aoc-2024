package day19

import (
	"fmt"
	"os"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day19/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 19 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected towels and designs")
	}
	towels := strings.Split(parts[0], ", ")
	var possible, ways int64
	for _, design := range strings.Split(parts[1], "\n") {
		memo := map[int]int64{len(design): 1}
		var count func(int) int64
		count = func(i int) int64 {
			if v, ok := memo[i]; ok {
				return v
			}
			var n int64
			for _, t := range towels {
				if strings.HasPrefix(design[i:], t) {
					n += count(i + len(t))
				}
			}
			memo[i] = n
			return n
		}
		n := count(0)
		if n > 0 {
			possible++
		}
		ways += n
	}
	return possible, ways, nil
}
