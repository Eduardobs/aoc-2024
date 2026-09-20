package day09

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = "2333133121414131402"

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 1928, 2858, err)
}
