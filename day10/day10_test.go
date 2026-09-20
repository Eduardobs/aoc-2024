package day10

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 36, 81, err)
}
