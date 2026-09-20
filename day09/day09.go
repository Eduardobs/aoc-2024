package day09

import (
	"fmt"
	"os"
	"strings"
)

type fileSpan struct {
	id, start, size int
}

type freeSpan struct {
	start, size int
	version     int
}

type freeCandidate struct {
	start, version, span int
}

type freeHeap []freeCandidate

func (h *freeHeap) push(candidate freeCandidate) {
	*h = append(*h, candidate)
	i := len(*h) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if (*h)[parent].start <= candidate.start {
			break
		}
		(*h)[i] = (*h)[parent]
		i = parent
	}
	(*h)[i] = candidate
}

func (h *freeHeap) discardTop() {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	if len(*h) == 0 {
		return
	}
	i := 0
	for {
		left := 2*i + 1
		if left >= len(*h) {
			break
		}
		child := left
		right := left + 1
		if right < len(*h) && (*h)[right].start < (*h)[left].start {
			child = right
		}
		if (*h)[child].start >= last.start {
			break
		}
		(*h)[i] = (*h)[child]
		i = child
	}
	(*h)[i] = last
}

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
	totalBlocks := 0
	for _, ch := range []byte(s) {
		if ch < '0' || ch > '9' {
			return 0, 0, fmt.Errorf("invalid digit")
		}
		totalBlocks += int(ch - '0')
	}
	disk := make([]int, 0, totalBlocks)
	files := make([]fileSpan, 0, (len(s)+1)/2)
	free := make([]freeSpan, 0, len(s)/2)
	position := 0
	for i, ch := range []byte(s) {
		n := int(ch - '0')
		v := -1
		if i%2 == 0 {
			v = i / 2
			files = append(files, fileSpan{id: v, start: position, size: n})
		} else if n > 0 {
			if len(free) > 0 && free[len(free)-1].start+free[len(free)-1].size == position {
				free[len(free)-1].size += n
			} else {
				free = append(free, freeSpan{start: position, size: n})
			}
		}
		for range n {
			disk = append(disk, v)
		}
		position += n
	}
	l, r := 0, len(disk)-1
	for l < r {
		for l < len(disk) && disk[l] >= 0 {
			l++
		}
		for r >= 0 && disk[r] < 0 {
			r--
		}
		if l < r {
			disk[l], disk[r] = disk[r], -1
		}
	}

	var gaps [10]freeHeap
	addFree := func(spanIndex int) {
		span := &free[spanIndex]
		for size := 1; size <= 9 && size <= span.size; size++ {
			gaps[size].push(freeCandidate{span.start, span.version, spanIndex})
		}
	}
	for i := range free {
		addFree(i)
	}
	for i := len(files) - 1; i >= 0; i-- {
		file := &files[i]
		if file.size == 0 {
			continue
		}
		candidates := &gaps[file.size]
		for len(*candidates) > 0 {
			candidate := (*candidates)[0]
			span := &free[candidate.span]
			if candidate.version == span.version && candidate.start == span.start && span.size >= file.size {
				break
			}
			candidates.discardTop()
		}
		if len(*candidates) == 0 || (*candidates)[0].start >= file.start {
			continue
		}
		spanIndex := (*candidates)[0].span
		span := &free[spanIndex]
		file.start = span.start
		span.start += file.size
		span.size -= file.size
		span.version++
		addFree(spanIndex)
	}
	var part2 int64
	for _, file := range files {
		positions := int64(file.size * (2*file.start + file.size - 1) / 2)
		part2 += int64(file.id) * positions
	}
	return checksum(disk), part2, nil
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
