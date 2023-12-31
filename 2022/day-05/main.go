package main

import (
	"errors"
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	// Part 1
	aoc.RunSolution(part1, "\n\n")

	// Part2
	aoc.RunSolution(part2, "\n\n")
}

func part1(input []string) (string, error) {
	if len(input) != 2 {
		return "", errors.New("wrong input")
	}

	s := newStacks(strings.Split(input[0], "\n"))
	return s.solve(input[1], s.moveN)
}

func part2(input []string) (string, error) {
	if len(input) != 2 {
		return "", errors.New("wrong input")
	}

	s := newStacks(strings.Split(input[0], "\n"))
	return s.solve(input[1], s.moveBulk)
}

type stacks []*crate

type crate struct {
	label string
	below *crate
}

// cratesInput:
// "    [D]    "
// "[N] [C]    "
// "[Z] [M] [P]"
// " 1   2   3 "
func newStacks(cratesInput []string) stacks {
	nCrates := len(strings.Fields(cratesInput[len(cratesInput)-1]))
	result := make(stacks, nCrates)

	// Go bottom to top
	for i := 0; i < nCrates; i++ {
		for lvl := len(cratesInput) - 2; lvl >= 0; lvl-- {
			label := cratesInput[lvl][1+i*4]
			// If we reached top of the stack `i`, then move onto the next one
			if label == ' ' {
				break
			}
			// Stack found crates
			result[i] = &crate{string(label), result[i]}
		}
	}

	return result
}

func (s stacks) moveBulk(from, to, n int) {
	bottomOnMove := s[from]
	// grab n crates
	for i := 0; i < n-1; i++ {
		bottomOnMove = bottomOnMove.below
	}

	topOnMove := s[from]
	s[from] = bottomOnMove.below
	bottomOnMove.below = s[to]
	s[to] = topOnMove
}

func (s stacks) moveN(from, to, n int) {
	for i := 0; i < n; i++ {
		s.move(from, to)
	}
}

func (s stacks) move(from, to int) {
	belowFrom := s[from].below

	crateOnMove := s[from]
	crateOnMove.below = s[to]

	s[from] = belowFrom
	s[to] = crateOnMove
}

func (s stacks) solve(input string, f func(int, int, int)) (string, error) {
	for _, line := range strings.Split(input, "\n") {
		from, to, n, err := parseLine(line)
		if err != nil {
			return "", err
		}

		f(from, to, n)
	}

	var result strings.Builder
	for _, top := range s {
		result.WriteString(top.label)
	}

	return result.String(), nil
}

func parseLine(line string) (int, int, int, error) {
	fields := strings.Fields(line)

	n, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, 0, 0, err
	}

	from, err := strconv.Atoi(fields[3])
	if err != nil {
		return 0, 0, 0, err
	}

	to, err := strconv.Atoi(fields[5])
	if err != nil {
		return 0, 0, 0, err
	}

	// Switch to 0-based indexing
	return from - 1, to - 1, n, nil
}
