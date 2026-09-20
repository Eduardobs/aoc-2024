package day11

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = "125 17"

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 55312, 65601038650482, err)
}
