package day22

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day22/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 22 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, int64, error) {
	var part1 int64
	totals := map[[4]int]int64{}
	for _, s := range strings.Fields(input) {
		secret, e := strconv.ParseInt(s, 10, 64)
		if e != nil {
			return 0, 0, e
		}
		prices := make([]int, 2001)
		prices[0] = int(secret % 10)
		for i := 1; i <= 2000; i++ {
			secret = next(secret)
			prices[i] = int(secret % 10)
		}
		part1 += secret
		seen := map[[4]int]bool{}
		for i := 4; i <= 2000; i++ {
			seq := [4]int{prices[i-3] - prices[i-4], prices[i-2] - prices[i-3], prices[i-1] - prices[i-2], prices[i] - prices[i-1]}
			if !seen[seq] {
				seen[seq] = true
				totals[seq] += int64(prices[i])
			}
		}
	}
	var part2 int64
	for _, n := range totals {
		if n > part2 {
			part2 = n
		}
	}
	return part1, part2, nil
}

func next(n int64) int64 {
	n = (n ^ (n * 64)) % 16777216
	n = (n ^ (n / 32)) % 16777216
	return (n ^ (n * 2048)) % 16777216
}
