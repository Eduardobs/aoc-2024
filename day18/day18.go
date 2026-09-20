package day18

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ x, y int }

var dirs = [...]point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func FirstProblem() (int64, error)   { a, _, e := solveFile(); return a, e }
func SecondProblem() (string, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, string, error) {
	b, e := os.ReadFile("./day18/input.txt")
	if e != nil {
		return 0, "", fmt.Errorf("read day 18 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, string, error) { return SolveWithSize(input, 71, 1024) }
func SolveWithSize(input string, size, first int) (int64, string, error) {
	bytes := []point{}
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var p point
		if _, e := fmt.Sscanf(line, "%d,%d", &p.x, &p.y); e != nil {
			return 0, "", fmt.Errorf("line %d: %w", i+1, e)
		}
		bytes = append(bytes, p)
	}
	if first > len(bytes) {
		first = len(bytes)
	}
	p1 := shortest(bytes[:first], size)
	if p1 < 0 {
		return 0, "", fmt.Errorf("path blocked by initial bytes")
	}
	lo, hi := first, len(bytes)
	for lo < hi {
		mid := (lo + hi) / 2
		if shortest(bytes[:mid+1], size) < 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo >= len(bytes) || shortest(bytes[:lo+1], size) >= 0 {
		return int64(p1), "", fmt.Errorf("path never becomes blocked")
	}
	return int64(p1), fmt.Sprintf("%d,%d", bytes[lo].x, bytes[lo].y), nil
}
func shortest(bs []point, size int) int {
	blocked := map[point]bool{}
	for _, p := range bs {
		blocked[p] = true
	}
	start, end := point{0, 0}, point{size - 1, size - 1}
	if blocked[start] || blocked[end] {
		return -1
	}
	dist := map[point]int{start: 0}
	q := []point{start}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p == end {
			return dist[p]
		}
		for _, d := range dirs {
			n := point{p.x + d.x, p.y + d.y}
			if n.x >= 0 && n.x < size && n.y >= 0 && n.y < size && !blocked[n] {
				if _, ok := dist[n]; !ok {
					dist[n] = dist[p] + 1
					q = append(q, n)
				}
			}
		}
	}
	return -1
}
