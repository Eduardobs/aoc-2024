package day01

import "testing"

const sample = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSample(t *testing.T) {
	a, b, err := parseLists(sample)
	if err != nil {
		t.Fatal(err)
	}
	if len(a) != 6 || len(b) != 6 {
		t.Fatalf("got list lengths %d and %d, want 6 and 6", len(a), len(b))
	}

	if got := totalDistance(append([]int(nil), a...), append([]int(nil), b...)); got != 11 {
		t.Fatalf("total distance = %d, want 11", got)
	}
	if got := similarityScore(a, b); got != 31 {
		t.Fatalf("similarity = %d, want 31", got)
	}
}

func TestMalformedInput(t *testing.T) {
	if _, _, err := parseLists("1 2 3"); err == nil {
		t.Fatal("expected malformed input error")
	}
}
