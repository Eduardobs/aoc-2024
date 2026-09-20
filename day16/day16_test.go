package day16

import (
	"testing"

	testutil "aoc-2024/tests"
)

const sample = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

func TestSample(t *testing.T) {
	part1, part2, err := Solve(sample)
	testutil.CheckNumbers(t, part1, part2, 7036, 45, err)
}
