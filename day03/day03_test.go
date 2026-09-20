package day03

import "testing"

func TestSamples(t *testing.T) {
	part1 := solve(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`, false)
	if part1 != 161 {
		t.Fatalf("part 1 = %d, want 161", part1)
	}

	part2 := solve(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`, true)
	if part2 != 48 {
		t.Fatalf("part 2 = %d, want 48", part2)
	}
}
