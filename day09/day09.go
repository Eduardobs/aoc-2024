package day09

import (
	"fmt"
	"os"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day09/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 9 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	s := strings.TrimSpace(input)
	disk := []int{}
	for i, ch := range []byte(s) {
		if ch < '0' || ch > '9' {
			return 0, 0, fmt.Errorf("invalid digit")
		}
		n := int(ch - '0')
		v := -1
		if i%2 == 0 {
			v = i / 2
		}
		for range n {
			disk = append(disk, v)
		}
	}
	a := append([]int(nil), disk...)
	l, r := 0, len(a)-1
	for l < r {
		for l < len(a) && a[l] >= 0 {
			l++
		}
		for r >= 0 && a[r] < 0 {
			r--
		}
		if l < r {
			a[l], a[r] = a[r], -1
		}
	}
	b := append([]int(nil), disk...)
	maxID := (len(s) - 1) / 2
	for id := maxID; id >= 0; id-- {
		start, size := -1, 0
		for i, v := range b {
			if v == id {
				if start < 0 {
					start = i
				}
				size++
			}
		}
		gapStart, gap := 0, 0
		dest := -1
		for i := 0; i < start; i++ {
			if b[i] < 0 {
				if gap == 0 {
					gapStart = i
				}
				gap++
				if gap >= size {
					dest = gapStart
					break
				}
			} else {
				gap = 0
			}
		}
		if dest >= 0 {
			for i := 0; i < size; i++ {
				b[dest+i] = id
				b[start+i] = -1
			}
		}
	}
	return checksum(a), checksum(b), nil
}
func checksum(d []int) int64 {
	var n int64
	for i, v := range d {
		if v >= 0 {
			n += int64(i * v)
		}
	}
	return n
}
