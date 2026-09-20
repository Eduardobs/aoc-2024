package day17

import "testing"

const sample = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestSample(t *testing.T) {
	output, registerA, err := Solve(sample)
	if err != nil {
		t.Fatal(err)
	}
	if output != "5,7,3,0" || registerA != 117440 {
		t.Fatalf("got (%q, %d), want (%q, %d)", output, registerA, "5,7,3,0", 117440)
	}
}
