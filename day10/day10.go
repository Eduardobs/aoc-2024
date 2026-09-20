package day10

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
	b, e := os.ReadFile("./day10/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 10 input: %w", e)
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
	var score, rating int64
	memo := map[point]int64{}
	var paths func(point) int64
	paths = func(p point) int64 {
		if v, ok := memo[p]; ok {
			return v
		}
		if g[p.r][p.c] == '9' {
			return 1
		}
		var n int64
		for _, d := range dirs {
			q := point{p.r + d.r, p.c + d.c}
			if q.r >= 0 && q.r < len(g) && q.c >= 0 && q.c < len(g[0]) && g[q.r][q.c] == g[p.r][p.c]+1 {
				n += paths(q)
			}
		}
		memo[p] = n
		return n
	}
	for r := range g {
		for c := range g[r] {
			if g[r][c] != '0' {
				continue
			}
			start := point{r, c}
			rating += paths(start)
			seen := map[point]bool{start: true}
			queue := []point{start}
			peaks := map[point]bool{}
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if g[p.r][p.c] == '9' {
					peaks[p] = true
				}
				for _, d := range dirs {
					q := point{p.r + d.r, p.c + d.c}
					if q.r >= 0 && q.r < len(g) && q.c >= 0 && q.c < len(g[0]) && !seen[q] && g[q.r][q.c] == g[p.r][p.c]+1 {
						seen[q] = true
						queue = append(queue, q)
					}
				}
			}
			score += int64(len(peaks))
		}
	}
	return score, rating, nil
}
