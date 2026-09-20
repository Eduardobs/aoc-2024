package day19

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 6, 16, err)
}
