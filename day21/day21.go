package day21

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct{ x, y int }
type memoKey struct {
	from, to byte
	depth    int
}

var numeric = map[byte]point{'7': {0, 0}, '8': {1, 0}, '9': {2, 0}, '4': {0, 1}, '5': {1, 1}, '6': {2, 1}, '1': {0, 2}, '2': {1, 2}, '3': {2, 2}, '0': {1, 3}, 'A': {2, 3}}
var directional = map[byte]point{'^': {1, 0}, 'A': {2, 0}, '<': {0, 1}, 'v': {1, 1}, '>': {2, 1}}

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day21/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 21 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	codes := strings.Fields(input)
	a, e := complexity(codes, 2)
	if e != nil {
		return 0, 0, e
	}
	b, e := complexity(codes, 25)
	return a, b, e
}
func complexity(codes []string, robots int) (int64, error) {
	memo := map[memoKey]int64{}
	var total int64
	for _, code := range codes {
		value, e := strconv.ParseInt(strings.TrimSuffix(code, "A"), 10, 64)
		if e != nil {
			return 0, e
		}
		var presses int64
		from := byte('A')
		for i := range code {
			best := int64(^uint64(0) >> 1)
			for _, path := range paths(numeric, from, code[i]) {
				cost := sequenceCost(path, robots, memo)
				if cost < best {
					best = cost
				}
			}
			presses += best
			from = code[i]
		}
		total += value * presses
	}
	return total, nil
}
func sequenceCost(seq string, depth int, memo map[memoKey]int64) int64 {
	if depth == 0 {
		return int64(len(seq))
	}
	var total int64
	from := byte('A')
	for i := range seq {
		k := memoKey{from, seq[i], depth}
		best, ok := memo[k]
		if !ok {
			best = int64(^uint64(0) >> 1)
			for _, p := range paths(directional, from, seq[i]) {
				v := sequenceCost(p, depth-1, memo)
				if v < best {
					best = v
				}
			}
			memo[k] = best
		}
		total += best
		from = seq[i]
	}
	return total
}
func paths(keypad map[byte]point, from, to byte) []string {
	a, b := keypad[from], keypad[to]
	valid := map[point]bool{}
	for _, p := range keypad {
		valid[p] = true
	}
	type step struct {
		ch byte
		d  point
	}
	moves := []step{{'<', point{-1, 0}}, {'>', point{1, 0}}, {'^', point{0, -1}}, {'v', point{0, 1}}}
	distance := abs(b.x-a.x) + abs(b.y-a.y)
	result := []string{}
	var build func(point, []byte)
	build = func(p point, route []byte) {
		if len(route) == distance {
			if p == b {
				result = append(result, string(route)+"A")
			}
			return
		}
		for _, move := range moves {
			next := point{p.x + move.d.x, p.y + move.d.y}
			if valid[next] && abs(b.x-next.x)+abs(b.y-next.y) == distance-len(route)-1 {
				build(next, append(route, move.ch))
			}
		}
	}
	build(a, nil)
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
