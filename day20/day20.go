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
	rows, cols := len(g), len(g[0])
	ds, de := distances(g, start), distances(g, end)
	normal := ds[end.r*cols+end.c]
	if normal < 0 {
		return 0, 0, fmt.Errorf("no path")
	}
	count := func(limit int) int64 {
		type offset struct{ dr, dc, distance int }
		offsets := make([]offset, 0, 2*limit*(limit+1))
		for dr := -limit; dr <= limit; dr++ {
			remain := limit - abs(dr)
			for dc := -remain; dc <= remain; dc++ {
				distance := abs(dr) + abs(dc)
				if distance >= 2 {
					offsets = append(offsets, offset{dr, dc, distance})
				}
			}
		}
		var n int64
		for cell, da := range ds {
			if da < 0 {
				continue
			}
			r, c := cell/cols, cell%cols
			for _, offset := range offsets {
				br, bc := r+offset.dr, c+offset.dc
				if br < 0 || br >= rows || bc < 0 || bc >= cols {
					continue
				}
				db := de[br*cols+bc]
				if db >= 0 && normal-(da+offset.distance+db) >= saving {
					n++
				}
			}
		}
		return n
	}
	return count(2), count(20), nil
}
func distances(g []string, start point) []int {
	rows, cols := len(g), len(g[0])
	d := make([]int, rows*cols)
	for i := range d {
		d[i] = -1
	}
	startIndex := start.r*cols + start.c
	d[startIndex] = 0
	q := make([]int, 1, rows*cols)
	q[0] = startIndex
	for head := 0; head < len(q); head++ {
		cell := q[head]
		p := point{cell / cols, cell % cols}
		for _, v := range dirs {
			n := point{p.r + v.r, p.c + v.c}
			if n.r < 0 || n.r >= rows || n.c < 0 || n.c >= cols || g[n.r][n.c] == '#' {
				continue
			}
			next := n.r*cols + n.c
			if d[next] < 0 {
				d[next] = d[cell] + 1
				q = append(q, next)
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
