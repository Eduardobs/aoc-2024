package main

import (
	"aoc-2024/day01"
	"aoc-2024/day02"
	"aoc-2024/day03"
	"aoc-2024/day04"
	"aoc-2024/day05"
	"aoc-2024/day06"
	"aoc-2024/day07"
	"aoc-2024/day08"
	"aoc-2024/day09"
	"aoc-2024/day10"
	"aoc-2024/day11"
	"aoc-2024/day12"
	"aoc-2024/day13"
	"aoc-2024/day14"
	"aoc-2024/day15"
	"aoc-2024/day16"
	"aoc-2024/day17"
	"aoc-2024/day18"
	"aoc-2024/day19"
	"aoc-2024/day20"
	"aoc-2024/day21"
	"aoc-2024/day22"
	"aoc-2024/day23"
	"aoc-2024/day24"
	"aoc-2024/day25"
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
	Execute(day05.FirstProblem, "Day 5, problem 1 result")
	Execute(day05.SecondProblem, "Day 5, problem 2 result")
	Execute(day06.FirstProblem, "Day 6, problem 1 result")
	Execute(day06.SecondProblem, "Day 6, problem 2 result")
	Execute(day07.FirstProblem, "Day 7, problem 1 result")
	Execute(day07.SecondProblem, "Day 7, problem 2 result")
	Execute(day08.FirstProblem, "Day 8, problem 1 result")
	Execute(day08.SecondProblem, "Day 8, problem 2 result")
	Execute(day09.FirstProblem, "Day 9, problem 1 result")
	Execute(day09.SecondProblem, "Day 9, problem 2 result")
	Execute(day10.FirstProblem, "Day 10, problem 1 result")
	Execute(day10.SecondProblem, "Day 10, problem 2 result")
	Execute(day11.FirstProblem, "Day 11, problem 1 result")
	Execute(day11.SecondProblem, "Day 11, problem 2 result")
	Execute(day12.FirstProblem, "Day 12, problem 1 result")
	Execute(day12.SecondProblem, "Day 12, problem 2 result")
	Execute(day13.FirstProblem, "Day 13, problem 1 result")
	Execute(day13.SecondProblem, "Day 13, problem 2 result")
	Execute(day14.FirstProblem, "Day 14, problem 1 result")
	Execute(day14.SecondProblem, "Day 14, problem 2 result")
	Execute(day15.FirstProblem, "Day 15, problem 1 result")
	Execute(day15.SecondProblem, "Day 15, problem 2 result")
	Execute(day16.FirstProblem, "Day 16, problem 1 result")
	Execute(day16.SecondProblem, "Day 16, problem 2 result")
	ExecuteText(day17.FirstProblem, "Day 17, problem 1 result")
	Execute(day17.SecondProblem, "Day 17, problem 2 result")
	Execute(day18.FirstProblem, "Day 18, problem 1 result")
	ExecuteText(day18.SecondProblem, "Day 18, problem 2 result")
	Execute(day19.FirstProblem, "Day 19, problem 1 result")
	Execute(day19.SecondProblem, "Day 19, problem 2 result")
	Execute(day20.FirstProblem, "Day 20, problem 1 result")
	Execute(day20.SecondProblem, "Day 20, problem 2 result")
	Execute(day21.FirstProblem, "Day 21, problem 1 result")
	Execute(day21.SecondProblem, "Day 21, problem 2 result")
	Execute(day22.FirstProblem, "Day 22, problem 1 result")
	Execute(day22.SecondProblem, "Day 22, problem 2 result")
	Execute(day23.FirstProblem, "Day 23, problem 1 result")
	ExecuteText(day23.SecondProblem, "Day 23, problem 2 result")
	Execute(day24.FirstProblem, "Day 24, problem 1 result")
	ExecuteText(day24.SecondProblem, "Day 24, problem 2 result")
	Execute(day25.FirstProblem, "Day 25 result")
}

func ExecuteText(funcExecute func() (string, error), successMsg string) {
	result, err := funcExecute()
	if err != nil {
		fmt.Printf("%s: error: %v\n", successMsg, err)
		return
	}
	fmt.Printf("%s: %s\n", successMsg, result)
}

func Execute(funcExecute func() (int64, error), successMsg string) {
	result, err := funcExecute()
	if err != nil {
		fmt.Printf("%s: error: %v\n", successMsg, err)
		return
	}

	fmt.Printf("%s: %d\n", successMsg, result)
}
