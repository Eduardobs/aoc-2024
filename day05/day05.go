package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }

func solveFile() (int64, int64, error) {
	b, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("read day 5 input: %w", err)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, int64, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return 0, 0, fmt.Errorf("expected rules and updates")
	}
	rules := map[int]map[int]bool{}
	for i, line := range strings.Split(sections[0], "\n") {
		var a, b int
		if _, err := fmt.Sscanf(strings.TrimSpace(line), "%d|%d", &a, &b); err != nil {
			return 0, 0, fmt.Errorf("rule %d: %w", i+1, err)
		}
		if rules[a] == nil {
			rules[a] = map[int]bool{}
		}
		rules[a][b] = true
	}
	var p1, p2 int64
	for i, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(strings.TrimSpace(line), ",")
		update := make([]int, len(parts))
		for j, s := range parts {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, 0, fmt.Errorf("update %d: %w", i+1, err)
			}
			update[j] = n
		}
		if valid(update, rules) {
			p1 += int64(update[len(update)/2])
			continue
		}
		fixed, err := reorder(update, rules)
		if err != nil {
			return 0, 0, fmt.Errorf("update %d: %w", i+1, err)
		}
		p2 += int64(fixed[len(fixed)/2])
	}
	return p1, p2, nil
}

func valid(update []int, rules map[int]map[int]bool) bool {
	pos := map[int]int{}
	for i, n := range update {
		pos[n] = i
	}
	for a, after := range rules {
		for b := range after {
			ia, oka := pos[a]
			ib, okb := pos[b]
			if oka && okb && ia > ib {
				return false
			}
		}
	}
	return true
}

func reorder(update []int, rules map[int]map[int]bool) ([]int, error) {
	present, indegree := map[int]bool{}, map[int]int{}
	for _, n := range update {
		present[n] = true
		indegree[n] = 0
	}
	for a, after := range rules {
		if !present[a] {
			continue
		}
		for b := range after {
			if present[b] {
				indegree[b]++
			}
		}
	}
	result := make([]int, 0, len(update))
	for len(result) < len(update) {
		found := false
		for _, n := range update {
			if present[n] && indegree[n] == 0 {
				found = true
				present[n] = false
				result = append(result, n)
				for b := range rules[n] {
					if present[b] {
						indegree[b]--
					}
				}
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("cyclic ordering rules")
		}
	}
	return result, nil
}
