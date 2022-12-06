package main

import (
	"errors"
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	// Part 1
	aoc.RunSolution(part1, "\n")

	// Part 2
	aoc.RunSolution(part2, "\n")
}

func part2(input []string) (string, error) {
	return solution(input, 14)
}

func part1(input []string) (string, error) {
	return solution(input, 4)
}

func solution(input []string, peekSize int) (string, error) {
	result := make([]string, len(input))
	for i, line := range input {
		n, err := newScanner(line).scan(peekSize)
		if err != nil {
			return "", err
		}

		result[i] = strconv.Itoa(n)
	}

	return strings.Join(result, ","), nil
}

type scanner struct {
	input    string
	position int
}

func newScanner(input string) *scanner {
	return &scanner{input: input}
}

func (s *scanner) peek(n int) byte {
	return s.input[s.position+n]
}

func (s *scanner) next() bool {
	s.position++
	return s.position <= len(s.input)
}

func (s *scanner) isMarker(peekSize int) bool {
	m := map[string]struct{}{}
	for i := 0; i < peekSize; i++ {
		b := s.peek(i)
		if _, ok := m[string(b)]; ok {
			return false // found duplicate
		}
		m[string(b)] = struct{}{}
	}
	return true
}

func (s *scanner) scan(peekSize int) (int, error) {
	for s.next() {
		if s.isMarker(peekSize) {
			return s.position + peekSize, nil
		}
	}

	return 0, errors.New("not found")
}
