package day14

import "testing"

const sample = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestSample(t *testing.T) {
	part1, _, err := SolveWithSize(sample, 11, 7)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 12 {
		t.Fatalf("part 1 = %d, want 12", part1)
	}
}

func TestPart2UsesFirstNonOverlappingArrangement(t *testing.T) {
	input := `p=0,0 v=0,0
p=1,0 v=-1,0`

	_, part2, err := SolveWithSize(input, 5, 5)
	if err != nil {
		t.Fatal(err)
	}
	if part2 != 0 {
		t.Fatalf("part 2 = %d, want 0", part2)
	}
}
