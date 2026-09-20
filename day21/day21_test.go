package day21

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = "029A\n980A\n179A\n456A\n379A"

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 126384, 154115708116294, err)
}
