package day24

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type gate struct{ a, op, b, out string }

func FirstProblem() (int64, error)   { a, _, e := solveFile(); return a, e }
func SecondProblem() (string, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, string, error) {
	b, e := os.ReadFile("./day24/input.txt")
	if e != nil {
		return 0, "", fmt.Errorf("read day 24 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, string, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("expected inputs and gates")
	}
	values := map[string]int{}
	for i, line := range strings.Split(parts[0], "\n") {
		p := strings.Split(line, ": ")
		if len(p) != 2 {
			return 0, "", fmt.Errorf("input line %d invalid", i+1)
		}
		n, e := strconv.Atoi(p[1])
		if e != nil {
			return 0, "", e
		}
		values[p[0]] = n
	}
	gates := []gate{}
	for i, line := range strings.Split(parts[1], "\n") {
		var g gate
		if _, e := fmt.Sscanf(line, "%s %s %s -> %s", &g.a, &g.op, &g.b, &g.out); e != nil {
			return 0, "", fmt.Errorf("gate %d: %w", i+1, e)
		}
		gates = append(gates, g)
	}
	pending := append([]gate(nil), gates...)
	for len(pending) > 0 {
		next := pending[:0]
		progress := false
		for _, g := range pending {
			a, oka := values[g.a]
			b, okb := values[g.b]
			if !oka || !okb {
				next = append(next, g)
				continue
			}
			switch g.op {
			case "AND":
				values[g.out] = a & b
			case "OR":
				values[g.out] = a | b
			case "XOR":
				values[g.out] = a ^ b
			default:
				return 0, "", fmt.Errorf("unknown operation %s", g.op)
			}
			progress = true
		}
		if !progress {
			return 0, "", fmt.Errorf("unresolvable circuit")
		}
		pending = append([]gate(nil), next...)
	}
	var part1 int64
	for name, v := range values {
		if strings.HasPrefix(name, "z") {
			bit, e := strconv.Atoi(name[1:])
			if e == nil && v == 1 {
				part1 |= int64(1) << bit
			}
		}
	}
	return part1, findSwaps(gates), nil
}

func findSwaps(gs []gate) string {
	maxZ := ""
	for _, g := range gs {
		if strings.HasPrefix(g.out, "z") && g.out > maxZ {
			maxZ = g.out
		}
	}
	consumers := map[string][]gate{}
	for _, g := range gs {
		consumers[g.a] = append(consumers[g.a], g)
		consumers[g.b] = append(consumers[g.b], g)
	}
	bad := map[string]bool{}
	isXY := func(s string) bool { return strings.HasPrefix(s, "x") || strings.HasPrefix(s, "y") }
	for _, g := range gs {
		if strings.HasPrefix(g.out, "z") && g.out != maxZ && g.op != "XOR" {
			bad[g.out] = true
		}
		if g.op == "XOR" && !strings.HasPrefix(g.out, "z") && !isXY(g.a) && !isXY(g.b) {
			bad[g.out] = true
		}
		if g.op == "AND" && !((g.a == "x00" && g.b == "y00") || (g.a == "y00" && g.b == "x00")) {
			ok := false
			for _, c := range consumers[g.out] {
				if c.op == "OR" {
					ok = true
				}
			}
			if !ok {
				bad[g.out] = true
			}
		}
		if g.op == "XOR" && isXY(g.a) && isXY(g.b) && g.a != "x00" && g.b != "x00" {
			ok := false
			for _, c := range consumers[g.out] {
				if c.op == "XOR" {
					ok = true
				}
			}
			if !ok {
				bad[g.out] = true
			}
		}
	}
	names := make([]string, 0, len(bad))
	for n := range bad {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}
