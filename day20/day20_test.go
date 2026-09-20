package day20

import "testing"

const sample = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestSample(t *testing.T) {
	part1, _, err := SolveWithSaving(sample, 2)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 44 {
		t.Fatalf("part 1 = %d, want 44", part1)
	}

	_, part2, err := SolveWithSaving(sample, 50)
	if err != nil {
		t.Fatal(err)
	}
	if part2 != 285 {
		t.Fatalf("part 2 = %d, want 285", part2)
	}
}
