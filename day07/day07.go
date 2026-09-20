package day07

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FirstProblem() (int64, error)  { a, _, e := solveFile(); return a, e }
func SecondProblem() (int64, error) { _, b, e := solveFile(); return b, e }
func solveFile() (int64, int64, error) {
	b, e := os.ReadFile("./day07/input.txt")
	if e != nil {
		return 0, 0, fmt.Errorf("read day 7 input: %w", e)
	}
	return Solve(string(b))
}

func Solve(input string) (int64, int64, error) {
	var p1, p2 int64
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("line %d: invalid equation", i+1)
		}
		target, e := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if e != nil {
			return 0, 0, e
		}
		fs := strings.Fields(parts[1])
		nums := make([]int64, len(fs))
		for j, s := range fs {
			nums[j], e = strconv.ParseInt(s, 10, 64)
			if e != nil {
				return 0, 0, e
			}
		}
		part1Possible := possible(target, nums, false)
		if part1Possible {
			p1 += target
		}
		if part1Possible || possible(target, nums, true) {
			p2 += target
		}
	}
	return p1, p2, nil
}

func possible(target int64, nums []int64, concat bool) bool {
	if len(nums) == 0 {
		return false
	}
	if target < 0 {
		return possibleForward(target, nums, concat)
	}
	for _, n := range nums {
		if n < 0 {
			return possibleForward(target, nums, concat)
		}
	}

	type state struct {
		i int
		v int64
	}
	failed := map[state]bool{}
	var search func(int, int64) bool
	search = func(i int, wanted int64) bool {
		if i == 0 {
			return wanted == nums[0]
		}
		current := state{i, wanted}
		if failed[current] {
			return false
		}
		n := nums[i]
		if wanted >= n && search(i-1, wanted-n) {
			return true
		}
		if n == 0 {
			if wanted == 0 {
				return true
			}
		} else if wanted%n == 0 && search(i-1, wanted/n) {
			return true
		}
		if concat && wanted >= n {
			power := int64(10)
			for x := n; x >= 10; x /= 10 {
				power *= 10
			}
			if wanted%power == n && search(i-1, wanted/power) {
				return true
			}
		}
		failed[current] = true
		return false
	}
	return search(len(nums)-1, target)
}

func possibleForward(target int64, nums []int64, concat bool) bool {
	type state struct {
		i int
		v int64
	}
	failed := map[state]bool{}
	var search func(int, int64) bool
	search = func(i int, value int64) bool {
		if i == len(nums) {
			return value == target
		}
		current := state{i, value}
		if failed[current] {
			return false
		}
		n := nums[i]
		if search(i+1, value+n) || search(i+1, value*n) {
			return true
		}
		if concat {
			power := int64(10)
			for x := n; x >= 10; x /= 10 {
				power *= 10
			}
			if search(i+1, value*power+n) {
				return true
			}
		}
		failed[current] = true
		return false
	}
	return search(1, nums[0])
}
