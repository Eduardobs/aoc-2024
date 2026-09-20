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
	treeTime := -1
	period := lcm(w, h)
	for t := 0; t < period; t++ {
		occupied := make(map[int]struct{}, len(rs))
		allDistinct := true
		for _, r := range rs {
			x, y := mod(r.x+t*r.vx, w), mod(r.y+t*r.vy, h)
			position := y*w + x
			if _, exists := occupied[position]; exists {
				allDistinct = false
				break
			}
			occupied[position] = struct{}{}
		}
		if allDistinct {
			treeTime = t
			break
		}
	}
	if treeTime < 0 {
		return 0, 0, fmt.Errorf("robots never form a non-overlapping arrangement")
	}
	return p1, int64(treeTime), nil
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
