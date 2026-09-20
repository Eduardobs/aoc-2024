package day14

import (
	"fmt"
	"os"
	"strings"
)

type robot struct{ x, y, vx, vy int }

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day14/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 14 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) { return SolveWithSize(input, 101, 103) }
func SolveWithSize(input string, w, h int) (int64, int64, error) {
	var rs []robot
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var r robot
		if _, e := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy); e != nil {
			return 0, 0, fmt.Errorf("line %d: %w", i+1, e)
		}
		rs = append(rs, r)
	}
	counts := [4]int{}
	for _, r := range rs {
		x, y := mod(r.x+100*r.vx, w), mod(r.y+100*r.vy, h)
		if x == w/2 || y == h/2 {
			continue
		}
		q := 0
		if x > w/2 {
			q++
		}
		if y > h/2 {
			q += 2
		}
		counts[q]++
	}
	p1 := int64(counts[0] * counts[1] * counts[2] * counts[3])
	bestT, bestArea := 0, int64(^uint64(0)>>1)
	period := lcm(w, h)
	for t := 0; t < period; t++ {
		minx, maxx, miny, maxy := w, 0, h, 0
		for _, r := range rs {
			x, y := mod(r.x+t*r.vx, w), mod(r.y+t*r.vy, h)
			if x < minx {
				minx = x
			}
			if x > maxx {
				maxx = x
			}
			if y < miny {
				miny = y
			}
			if y > maxy {
				maxy = y
			}
		}
		area := int64((maxx - minx + 1) * (maxy - miny + 1))
		if area < bestArea {
			bestArea = area
			bestT = t
		}
	}
	return p1, int64(bestT), nil
}
func mod(a, n int) int {
	a %= n
	if a < 0 {
		a += n
	}
	return a
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int { return a / gcd(a, b) * b }
