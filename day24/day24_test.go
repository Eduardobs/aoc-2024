package day24

import "testing"

const sample = `x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`

func TestSample(t *testing.T) {
	part1, _, err := Solve(sample)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 4 {
		t.Fatalf("part 1 = %d, want 4", part1)
	}
}
