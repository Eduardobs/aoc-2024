package day15

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }

var move = map[byte]point{'^': {-1, 0}, 'v': {1, 0}, '<': {0, -1}, '>': {0, 1}}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day15/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 15 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected map and moves")
	}
	moves := strings.ReplaceAll(parts[1], "\n", "")
	g1 := toGrid(parts[0])
	p1 := run(g1, moves, false)
	lines := strings.Split(parts[0], "\n")
	wide := make([]string, len(lines))
	for i, s := range lines {
		var b strings.Builder
		for _, ch := range s {
			switch ch {
			case '#':
				b.WriteString("##")
			case '.':
				b.WriteString("..")
			case 'O':
				b.WriteString("[]")
			case '@':
				b.WriteString("@.")
			}
		}
		wide[i] = b.String()
	}
	p2 := run(toGrid(strings.Join(wide, "\n")), moves, true)
	return p1, p2, nil
}
func toGrid(s string) [][]byte {
	ls := strings.Split(s, "\n")
	g := make([][]byte, len(ls))
	for i := range ls {
		g[i] = []byte(ls[i])
	}
	return g
}
func run(g [][]byte, moves string, wide bool) int64 {
	var pos point
	for r := range g {
		for c := range g[r] {
			if g[r][c] == '@' {
				pos = point{r, c}
				g[r][c] = '.'
			}
		}
	}
	for _, ch := range []byte(moves) {
		d := move[ch]
		n := point{pos.r + d.r, pos.c + d.c}
		if g[n.r][n.c] == '#' {
			continue
		}
		if g[n.r][n.c] == '.' {
			pos = n
			continue
		}
		boxes := map[point]bool{}
		if collect(g, n, d, wide, boxes) {
			list := make([]point, 0, len(boxes))
			for p := range boxes {
				list = append(list, p)
			}
			for i := 0; i < len(list); i++ {
				for j := i + 1; j < len(list); j++ {
					if list[i].r*d.r+list[i].c*d.c < list[j].r*d.r+list[j].c*d.c {
						list[i], list[j] = list[j], list[i]
					}
				}
			}
			for _, p := range list {
				q := point{p.r + d.r, p.c + d.c}
				g[q.r][q.c] = g[p.r][p.c]
				g[p.r][p.c] = '.'
			}
			pos = n
		}
	}
	var sum int64
	for r := range g {
		for c, ch := range g[r] {
			if ch == 'O' || ch == '[' {
				sum += int64(100*r + c)
			}
		}
	}
	return sum
}
func collect(g [][]byte, p, d point, wide bool, boxes map[point]bool) bool {
	ch := g[p.r][p.c]
	if ch == '.' {
		return true
	}
	if ch == '#' {
		return false
	}
	if boxes[p] {
		return true
	}
	boxes[p] = true
	if wide && (ch == '[' || ch == ']') && d.r != 0 {
		other := point{p.r, p.c + 1}
		if ch == ']' {
			other.c = p.c - 1
		}
		if !collect(g, other, d, wide, boxes) {
			return false
		}
	}
	return collect(g, point{p.r + d.r, p.c + d.c}, d, wide, boxes)
}
