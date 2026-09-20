package day23

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func FirstProblem() (int64, error)   { a, _, e := solveFile(); return a, e }
func SecondProblem() (string, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, string, error) {
	b, e := os.ReadFile("./day23/input.txt")
	if e != nil {
		return 0, "", fmt.Errorf("read day 23 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, string, error) {
	g := map[string]map[string]bool{}
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		p := strings.Split(line, "-")
		if len(p) != 2 {
			return 0, "", fmt.Errorf("line %d: invalid edge", i+1)
		}
		if g[p[0]] == nil {
			g[p[0]] = map[string]bool{}
		}
		if g[p[1]] == nil {
			g[p[1]] = map[string]bool{}
		}
		g[p[0]][p[1]] = true
		g[p[1]][p[0]] = true
	}
	nodes := make([]string, 0, len(g))
	for n := range g {
		nodes = append(nodes, n)
	}
	sort.Strings(nodes)
	var triangles int64
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if !g[nodes[i]][nodes[j]] {
				continue
			}
			for k := j + 1; k < len(nodes); k++ {
				if g[nodes[i]][nodes[k]] && g[nodes[j]][nodes[k]] && (nodes[i][0] == 't' || nodes[j][0] == 't' || nodes[k][0] == 't') {
					triangles++
				}
			}
		}
	}
	best := []string{}
	p := map[string]bool{}
	for _, n := range nodes {
		p[n] = true
	}
	var bron func([]string, map[string]bool, map[string]bool)
	bron = func(r []string, p, x map[string]bool) {
		if len(p) == 0 && len(x) == 0 {
			if len(r) > len(best) {
				best = append([]string(nil), r...)
			}
			return
		}
		var pivot string
		max := -1
		for u := range union(p, x) {
			n := 0
			for v := range p {
				if g[u][v] {
					n++
				}
			}
			if n > max {
				max = n
				pivot = u
			}
		}
		candidates := difference(p, g[pivot])
		for _, v := range sortedKeys(candidates) {
			nr := append(append([]string(nil), r...), v)
			bron(nr, intersection(p, g[v]), intersection(x, g[v]))
			delete(p, v)
			x[v] = true
		}
	}
	bron(nil, p, map[string]bool{})
	sort.Strings(best)
	return triangles, strings.Join(best, ","), nil
}
func union(a, b map[string]bool) map[string]bool {
	r := map[string]bool{}
	for k := range a {
		r[k] = true
	}
	for k := range b {
		r[k] = true
	}
	return r
}
func difference(a, b map[string]bool) map[string]bool {
	r := map[string]bool{}
	for k := range a {
		if !b[k] {
			r[k] = true
		}
	}
	return r
}
func intersection(a, b map[string]bool) map[string]bool {
	r := map[string]bool{}
	for k := range a {
		if b[k] {
			r[k] = true
		}
	}
	return r
}
func sortedKeys(m map[string]bool) []string {
	r := make([]string, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	sort.Strings(r)
	return r
}
