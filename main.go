package main

import (
	"aoc-2024/day01"
	"aoc-2024/day02"
	"aoc-2024/day03"
	"aoc-2024/day04"
	"fmt"
)

func main() {
	Execute(day01.FirstProblem, "Day 1, problem 1 result")
	Execute(day01.SecondProblem, "Day 1, problem 2 result")
	Execute(day02.FirstProblem, "Day 2, problem 1 result")
	Execute(day02.SecondProblem, "Day 2, problem 2 result")
	Execute(day03.FirstProblem, "Day 3, problem 1 result")
	Execute(day03.SecondProblem, "Day 3, problem 2 result")
	Execute(day04.FirstProblem, "Day 4, problem 1 result")
	Execute(day04.SecondProblem, "Day 4, problem 2 result")
}

func Execute(funcExecute func() (int64, error), successMsg string) {
	result, err := funcExecute()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %d\n", successMsg, result)
}
