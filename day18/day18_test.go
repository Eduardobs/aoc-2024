package day18

import "testing"

const sample = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestSample(t *testing.T) {
	part1, part2, err := SolveWithSize(sample, 7, 12)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 22 || part2 != "6,1" {
		t.Fatalf("got (%d, %q), want (%d, %q)", part1, part2, 22, "6,1")
	}
}
