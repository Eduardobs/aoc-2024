package day12

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }

var dirs = [...]point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day12/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 12 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	g := strings.Split(strings.TrimSpace(input), "\n")
	if len(g) == 0 || len(g[0]) == 0 {
		return 0, 0, fmt.Errorf("empty grid")
	}
	for _, s := range g {
		if len(s) != len(g[0]) {
			return 0, 0, fmt.Errorf("ragged grid")
		}
	}
	seen := map[point]bool{}
	var p1, p2 int64
	inside := func(p point) bool { return p.r >= 0 && p.r < len(g) && p.c >= 0 && p.c < len(g[0]) }
	for r := range g {
		for c := range g[r] {
			start := point{r, c}
			if seen[start] {
				continue
			}
			ch := g[r][c]
			queue := []point{start}
			seen[start] = true
			region := map[point]bool{}
			perim := 0
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				region[p] = true
				for _, d := range dirs {
					q := point{p.r + d.r, p.c + d.c}
					if !inside(q) || g[q.r][q.c] != ch {
						perim++
						continue
					}
					if !seen[q] {
						seen[q] = true
						queue = append(queue, q)
					}
				}
			}
			corners := 0
			for p := range region {
				for _, pair := range [][2]point{{{-1, 0}, {0, -1}}, {{-1, 0}, {0, 1}}, {{1, 0}, {0, -1}}, {{1, 0}, {0, 1}}} {
					a := point{p.r + pair[0].r, p.c + pair[0].c}
					b := point{p.r + pair[1].r, p.c + pair[1].c}
					diag := point{p.r + pair[0].r + pair[1].r, p.c + pair[0].c + pair[1].c}
					if !region[a] && !region[b] || region[a] && region[b] && !region[diag] {
						corners++
					}
				}
			}
			area := len(region)
			p1 += int64(area * perim)
			p2 += int64(area * corners)
		}
	}
	return p1, p2, nil
}
