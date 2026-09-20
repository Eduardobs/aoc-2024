package day08

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day08/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 8 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	g := strings.Split(strings.TrimSpace(input), "\n")
	if len(g) == 0 || len(g[0]) == 0 {
		return 0, 0, fmt.Errorf("empty grid")
	}
	ants := map[byte][]point{}
	for r := range g {
		if len(g[r]) != len(g[0]) {
			return 0, 0, fmt.Errorf("ragged grid")
		}
		for c := range g[r] {
			if g[r][c] != '.' {
				ants[g[r][c]] = append(ants[g[r][c]], point{r, c})
			}
		}
	}
	in := func(p point) bool { return p.r >= 0 && p.r < len(g) && p.c >= 0 && p.c < len(g[0]) }
	a1, a2 := map[point]bool{}, map[point]bool{}
	for _, ps := range ants {
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				dr, dc := ps[j].r-ps[i].r, ps[j].c-ps[i].c
				p := point{ps[i].r - dr, ps[i].c - dc}
				if in(p) {
					a1[p] = true
				}
				p = point{ps[j].r + dr, ps[j].c + dc}
				if in(p) {
					a1[p] = true
				}
				p = ps[i]
				for in(p) {
					a2[p] = true
					p = point{p.r - dr, p.c - dc}
				}
				p = ps[j]
				for in(p) {
					a2[p] = true
					p = point{p.r + dr, p.c + dc}
				}
			}
		}
	}
	return int64(len(a1)), int64(len(a2)), nil
}
