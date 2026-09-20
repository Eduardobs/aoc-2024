package day02

import "testing"

const sample = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestSample(t *testing.T) {
	part1, err := solve(sample, false)
	if err != nil {
		t.Fatal(err)
	}
	part2, err := solve(sample, true)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 2 || part2 != 4 {
		t.Fatalf("got (%d, %d), want (2, 4)", part1, part2)
	}
}

func TestShortReportsDoNotPanic(t *testing.T) {
	if !safe(nil) || !safe([]int{1}) || !safeByRemoving([]int{1, 5}) {
		t.Fatal("reports with fewer than two remaining levels should be safe")
	}
}
