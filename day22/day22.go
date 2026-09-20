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
	const (
		changes      = 19
		sequenceKeys = changes * changes * changes * changes
		lastThree    = changes * changes * changes
	)
	var part1 int64
	totals := make([]int64, sequenceKeys)
	seen := make([]int, sequenceKeys)
	for buyer, s := range strings.Fields(input) {
		secret, e := strconv.ParseInt(s, 10, 64)
		if e != nil {
			return 0, 0, e
		}
		previousPrice := int(secret % 10)
		key := 0
		stamp := buyer + 1
		for i := 1; i <= 2000; i++ {
			secret = next(secret)
			price := int(secret % 10)
			key = (key%lastThree)*changes + price - previousPrice + 9
			previousPrice = price
			if i >= 4 && seen[key] != stamp {
				seen[key] = stamp
				totals[key] += int64(price)
			}
		}
		part1 += secret
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
	const mask = 1<<24 - 1
	n = (n ^ (n << 6)) & mask
	n = (n ^ (n >> 5)) & mask
	return (n ^ (n << 11)) & mask
}
