package day16

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type point struct{ r, c int }
type state struct {
	point
	d int
}

var dirs = [...]point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const inf = int64(^uint64(0) >> 1)

type item struct {
	s    state
	cost int64
}
type pq []item

func (p pq) Len() int               { return len(p) }
func (p pq) Less(i, j int) bool     { return p[i].cost < p[j].cost }
func (p pq) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)            { *p = append(*p, x.(item)) }
func (p *pq) Pop() any              { o := *p; x := o[len(o)-1]; *p = o[:len(o)-1]; return x }
func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day16/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 16 input: %w", e)
	}
	return Solve(string(b))
}
func Solve(input string) (int64, int64, error) {
	g := strings.Split(strings.TrimSpace(input), "\n")
	var start, end point
	for r := range g {
		for c, ch := range g[r] {
			if ch == 'S' {
				start = point{r, c}
			}
			if ch == 'E' {
				end = point{r, c}
			}
		}
	}
	from := dijkstra(g, []state{{start, 0}}, false)
	best := inf
	ends := []state{}
	for d := 0; d < 4; d++ {
		s := state{end, d}
		if from[s] < best {
			best = from[s]
			ends = []state{s}
		} else if from[s] == best {
			ends = append(ends, s)
		}
	}
	to := dijkstra(g, ends, true)
	tiles := map[point]bool{}
	for s, a := range from {
		if b, ok := to[s]; ok && a+b == best {
			tiles[s.point] = true
		}
	}
	if best == inf {
		return 0, 0, fmt.Errorf("no path")
	}
	return best, int64(len(tiles)), nil
}
func dijkstra(g []string, starts []state, reverse bool) map[state]int64 {
	dist := map[state]int64{}
	q := &pq{}
	heap.Init(q)
	for _, s := range starts {
		dist[s] = 0
		heap.Push(q, item{s, 0})
	}
	for q.Len() > 0 {
		it := heap.Pop(q).(item)
		if dist[it.s] != it.cost {
			continue
		}
		next := []item{{state{it.s.point, (it.s.d + 1) % 4}, it.cost + 1000}, {state{it.s.point, (it.s.d + 3) % 4}, it.cost + 1000}}
		step := dirs[it.s.d]
		if reverse {
			step = point{-step.r, -step.c}
		}
		p := point{it.s.r + step.r, it.s.c + step.c}
		if p.r >= 0 && p.r < len(g) && p.c >= 0 && p.c < len(g[0]) && g[p.r][p.c] != '#' {
			next = append(next, item{state{p, it.s.d}, it.cost + 1})
		}
		for _, n := range next {
			if old, ok := dist[n.s]; !ok || n.cost < old {
				dist[n.s] = n.cost
				heap.Push(q, n)
			}
		}
	}
	return dist
}
