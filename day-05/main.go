package main

import (
	"errors"
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	// Part 1
	aoc.RunSolution(getTops, "\n\n")
}

func getTops(input []string) (string, error) {
	if len(input) != 2 {
		return "", errors.New("wrong input")
	}

	s := newStacks(strings.Split(input[0], "\n"))

	for _, line := range strings.Split(input[1], "\n") {
		from, to, n, err := parseLine(line)
		if err != nil {
			return "", err
		}

		s.moveN(from, to, n)
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

	return from - 1, to - 1, n, nil
}

type stacks []*crate

func newStacks(cratesInput []string) stacks {
	nCrates := len(strings.Fields(cratesInput[len(cratesInput)-1]))
	result := make(stacks, nCrates)

	// cratesInput:
	// "    [D]    "
	// "[N] [C]    "
	// "[Z] [M] [P]"
	// " 1   2   3 "
	for i := len(cratesInput) - 2; i >= 0; i-- {
		for j := 0; j < nCrates; j++ {
			if label := cratesInput[i][1+j*4]; label != ' ' {
				result[j] = &crate{string(label), result[j]}
			}
		}
	}

	return result
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

type crate struct {
	label string
	below *crate
}
