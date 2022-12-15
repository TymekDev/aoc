package main

import (
	"strings"
)

type herd []*monkey

func (h herd) round() {
	for i, m := range h {
		_ = i
		_ = m
	}
}

func newHerdFromInput(input []string) (herd, error) {
	result := make(herd, len(input))
	for i, line := range input {
		m, err := newMonkeyFromInput(strings.Split(line, "\n"))
		if err != nil {
			return nil, err
		}

		result[i] = m
	}

	return result, nil
}

type monkey struct {
	operation func(int) (int, error)
	test      func(int) int
	items     []int
}

// newMonkeyFromInput expects an input slice with a following structure:
//
//	0: "Monkey 0:"
//	1: "  Starting items: 79, 98"
//	2: "  Operation: new = old * 19"
//	3: "  Test: divisible by 23"
//	4: "    If true: throw to monkey 2"
//	5: "    If false: throw to monkey 3"
func newMonkeyFromInput(input []string) (*monkey, error) {
	items, err := parseItems(input[1])
	if err != nil {
		return nil, err
	}

	operation, err := parseOperation(input[2])
	if err != nil {
		return nil, err
	}

	test, err := parseTest(input[3], input[4], input[5])
	if err != nil {
		return nil, err
	}

	return &monkey{
		items:     items,
		operation: operation,
		test:      test,
	}, nil
}
