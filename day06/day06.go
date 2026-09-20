package day06

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }
type state struct {
	point
	d int
}

var dirs = [...]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day06/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 6 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, int64, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0, 0, fmt.Errorf("empty grid")
	}
	start := point{-1, -1}
	for r := range lines {
		if len(lines[r]) != len(lines[0]) {
			return 0, 0, fmt.Errorf("ragged grid")
		}
		if c := strings.IndexByte(lines[r], '^'); c >= 0 {
			start = point{r, c}
		}
	}
	if start.r < 0 {
		return 0, 0, fmt.Errorf("guard not found")
	}
	visited, _ := walk(lines, start, point{-1, -1})
	var loops int64
	for p := range visited {
		if p != start && lines[p.r][p.c] == '.' {
			_, loop := walk(lines, start, p)
			if loop {
				loops++
			}
		}
	}
	return int64(len(visited)), loops, nil
}

func walk(grid []string, start, obstacle point) (map[point]bool, bool) {
	seen, visited := map[state]bool{}, map[point]bool{}
	s := state{start, 0}
	for {
		if seen[s] {
			return visited, true
		}
		seen[s] = true
		visited[s.point] = true
		n := point{s.r + dirs[s.d].r, s.c + dirs[s.d].c}
		if n.r < 0 || n.r >= len(grid) || n.c < 0 || n.c >= len(grid[0]) {
			return visited, false
		}
		if grid[n.r][n.c] == '#' || n == obstacle {
			s.d = (s.d + 1) % 4
		} else {
			s.point = n
		}
	}
}
