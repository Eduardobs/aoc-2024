package day25

import (
	"fmt"
	"os"
	"strings"
)

func FirstProblem() (int64, error) {
	b, e := os.ReadFile("./day25/input.txt")
	if e != nil {
		return 0, fmt.Errorf("read day 25 input: %w", e)
	}
	return Solve(string(b))
}
func SecondProblem() (int64, error) { return 0, nil }
func Solve(input string) (int64, error) {
	locks, keys := [][]int{}, [][]int{}
	width := -1
	for i, block := range strings.Split(strings.TrimSpace(input), "\n\n") {
		lines := strings.Split(block, "\n")
		if len(lines) != 7 {
			return 0, fmt.Errorf("schematic %d: expected 7 rows", i+1)
		}
		if width < 0 {
			width = len(lines[0])
		} else if len(lines[0]) != width {
			return 0, fmt.Errorf("schematic %d: expected width %d", i+1, width)
		}
		h := make([]int, len(lines[0]))
		for _, line := range lines {
			if len(line) != len(h) {
				return 0, fmt.Errorf("schematic %d is ragged", i+1)
			}
			for c, ch := range line {
				if ch == '#' {
					h[c]++
				}
			}
		}
		if strings.Count(lines[0], "#") == len(h) {
			locks = append(locks, h)
		} else {
			keys = append(keys, h)
		}
	}
	var fit int64
	for _, l := range locks {
		for _, k := range keys {
			ok := true
			for c := range l {
				if l[c]+k[c] > 7 {
					ok = false
					break
				}
			}
			if ok {
				fit++
			}
		}
	}
	return fit, nil
}
