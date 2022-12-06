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
}

func part1(input []string) (string, error) {
	result := make([]string, len(input))
	for i, line := range input {
		n, err := newScanner(line).scan()
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

func (s *scanner) isMarker() bool {
	m := map[string]struct{}{}
	for i := 0; i < 4; i++ {
		b := s.peek(i)
		if _, ok := m[string(b)]; ok {
			return false // found duplicate
		}
		m[string(b)] = struct{}{}
	}
	return true
}

func (s *scanner) scan() (int, error) {
	for s.next() {
		if s.isMarker() {
			// 1-base indexed end of marker
			return s.position + 4, nil
		}
	}

	return 0, errors.New("not found")
}
