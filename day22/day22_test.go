package day22

import "testing"

func TestPart1Sample(t *testing.T) {
	part1, _, err := Solve("1\n10\n100\n2024")
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 37327623 {
		t.Fatalf("part 1 = %d, want 37327623", part1)
	}
}

func TestPart2Sample(t *testing.T) {
	_, part2, err := Solve("1\n2\n3\n2024")
	if err != nil {
		t.Fatal(err)
	}
	if part2 != 23 {
		t.Fatalf("part 2 = %d, want 23", part2)
	}
}
