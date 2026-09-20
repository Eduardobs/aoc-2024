package day17

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstProblem() (string, error) { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (string, int64, error) {
	b, e := os.ReadFile("./day17/input.txt")
	if e != nil {
		return "", 0, fmt.Errorf("read day 17 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (string, int64, error) {
	var a, b, c int64
	var raw string
	if _, e := fmt.Sscanf(strings.TrimSpace(input), "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s", &a, &b, &c, &raw); e != nil {
		return "", 0, e
	}
	parts := strings.Split(raw, ",")
	program := make([]int, len(parts))
	for i, s := range parts {
		n, e := strconv.Atoi(s)
		if e != nil {
			return "", 0, e
		}
		program[i] = n
	}
	out, e := run(program, a, b, c)
	if e != nil {
		return "", 0, e
	}
	text := make([]string, len(out))
	for i, n := range out {
		text[i] = strconv.Itoa(n)
	}
	candidates := []int64{0}
	for i := len(program) - 1; i >= 0; i-- {
		next := []int64{}
		for _, base := range candidates {
			for digit := int64(0); digit < 8; digit++ {
				candidate := base*8 + digit
				o, e := run(program, candidate, b, c)
				if e == nil && equal(o, program[i:]) {
					next = append(next, candidate)
				}
			}
		}
		if len(next) == 0 {
			return strings.Join(text, ","), 0, fmt.Errorf("no quine register value")
		}
		candidates = next
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	return strings.Join(text, ","), candidates[0], nil
}

func run(program []int, a, b, c int64) ([]int, error) {
	ip := 0
	out := []int{}
	combo := func(x int) (int64, error) {
		switch x {
		case 0, 1, 2, 3:
			return int64(x), nil
		case 4:
			return a, nil
		case 5:
			return b, nil
		case 6:
			return c, nil
		}
		return 0, fmt.Errorf("invalid combo operand %d", x)
	}
	for ip >= 0 && ip < len(program) {
		if ip+1 >= len(program) {
			return nil, fmt.Errorf("truncated instruction")
		}
		op, arg := program[ip], program[ip+1]
		cv, e := combo(arg)
		switch op {
		case 0:
			if e != nil {
				return nil, e
			}
			a = dividePow2(a, cv)
		case 1:
			b ^= int64(arg)
		case 2:
			if e != nil {
				return nil, e
			}
			b = cv % 8
		case 3:
			if a != 0 {
				ip = arg
				continue
			}
		case 4:
			b ^= c
		case 5:
			if e != nil {
				return nil, e
			}
			out = append(out, int(cv%8))
		case 6:
			if e != nil {
				return nil, e
			}
			b = dividePow2(a, cv)
		case 7:
			if e != nil {
				return nil, e
			}
			c = dividePow2(a, cv)
		default:
			return nil, fmt.Errorf("invalid opcode %d", op)
		}
		ip += 2
	}
	return out, nil
}
func dividePow2(n, p int64) int64 {
	if p >= 63 {
		return 0
	}
	return n / (int64(1) << p)
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
