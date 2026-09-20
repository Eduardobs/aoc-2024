package day15

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = "#####\n#@O.#\n#####\n\n>>"

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 103, 105, err)
}
