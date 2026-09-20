package day25

import "testing"

const sample = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func TestSample(t *testing.T) {
	part1, err := Solve(sample)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 3 {
		t.Fatalf("part 1 = %d, want 3", part1)
	}
}
