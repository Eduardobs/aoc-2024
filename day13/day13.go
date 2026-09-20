package day13

import (
	"fmt"
	"os"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day13/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 13 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	var p1, p2 int64
	for i, block := range strings.Split(strings.TrimSpace(input), "\n\n") {
		var ax, ay, bx, by, px, py int64
		lines := strings.Split(block, "\n")
		if len(lines) != 3 {
			return 0, 0, fmt.Errorf("machine %d invalid", i+1)
		}
		if _, e := fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &ax, &ay); e != nil {
			return 0, 0, e
		}
		if _, e := fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bx, &by); e != nil {
			return 0, 0, e
		}
		if _, e := fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py); e != nil {
			return 0, 0, e
		}
		p1 += tokens(ax, ay, bx, by, px, py, 100)
		p2 += tokens(ax, ay, bx, by, px+10000000000000, py+10000000000000, -1)
	}
	return p1, p2, nil
}
func tokens(ax, ay, bx, by, px, py, limit int64) int64 {
	det := ax*by - ay*bx
	if det == 0 {
		return 0
	}
	an, bn := px*by-py*bx, ax*py-ay*px
	if an%det != 0 || bn%det != 0 {
		return 0
	}
	a, b := an/det, bn/det
	if a < 0 || b < 0 || limit >= 0 && (a > limit || b > limit) {
		return 0
	}
	return 3*a + b
}
