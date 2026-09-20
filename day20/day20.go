package day20

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }

var dirs = [...]point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day20/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 20 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) { return SolveWithSaving(input, 100) }
func SolveWithSaving(input string, saving int) (int64, int64, error) {
	g := strings.Split(strings.TrimSpace(input), "\n")
	var start, end point
	for r := range g {
		for c, ch := range g[r] {
			if ch == 'S' {
				start = point{r, c}
			}
			if ch == 'E' {
				end = point{r, c}
			}
		}
	}
	ds, de := distances(g, start), distances(g, end)
	normal, ok := ds[end]
	if !ok {
		return 0, 0, fmt.Errorf("no path")
	}
	count := func(limit int) int64 {
		var n int64
		for a, da := range ds {
			for dr := -limit; dr <= limit; dr++ {
				remain := limit - abs(dr)
				for dc := -remain; dc <= remain; dc++ {
					length := abs(dr) + abs(dc)
					if length < 2 {
						continue
					}
					b := point{a.r + dr, a.c + dc}
					if db, ok := de[b]; ok && normal-(da+length+db) >= saving {
						n++
					}
				}
			}
		}
		return n
	}
	return count(2), count(20), nil
}
func distances(g []string, start point) map[point]int {
	d := map[point]int{start: 0}
	q := []point{start}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, v := range dirs {
			n := point{p.r + v.r, p.c + v.c}
			if n.r >= 0 && n.r < len(g) && n.c >= 0 && n.c < len(g[0]) && g[n.r][n.c] != '#' {
				if _, ok := d[n]; !ok {
					d[n] = d[p] + 1
					q = append(q, n)
				}
			}
		}
	}
	return d
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
