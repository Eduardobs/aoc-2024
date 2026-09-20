package day04

import "testing"

const sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestSample(t *testing.T) {
	grid, err := parseGrid(sample)
	if err != nil {
		t.Fatal(err)
	}
	if got := countXMAS(grid); got != 18 {
		t.Fatalf("part 1 = %d, want 18", got)
	}
	if got := countXMASCrosses(grid); got != 9 {
		t.Fatalf("part 2 = %d, want 9", got)
	}
}

func TestCrossOrientations(t *testing.T) {
	for _, grid := range []string{
		"M.S\n.A.\nM.S",
		"S.M\n.A.\nS.M",
		"M.M\n.A.\nS.S",
		"S.S\n.A.\nM.M",
	} {
		lines, err := parseGrid(grid)
		if err != nil {
			t.Fatal(err)
		}
		if got := countXMASCrosses(lines); got != 1 {
			t.Errorf("countXMASCrosses(%q) = %d, want 1", grid, got)
		}
	}
}

func TestRejectsRaggedGrid(t *testing.T) {
	if _, err := parseGrid("XMAS\nXMA"); err == nil {
		t.Fatal("expected ragged-grid error")
	}
}
