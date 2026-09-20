package day23

import "testing"

const sample = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestSample(t *testing.T) {
	part1, password, err := Solve(sample)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 7 || password != "co,de,ka,ta" {
		t.Fatalf("got (%d, %q), want (%d, %q)", part1, password, 7, "co,de,ka,ta")
	}
}
