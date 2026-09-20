package day12

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = `AAAA
BBCD
BBCC
EEEC`

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 140, 80, err)
}
