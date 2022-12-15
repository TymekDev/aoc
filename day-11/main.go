package main

import (
	"fmt"
	"sort"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

const (
	_verbose      = false
	_roundSummary = true
)

func main() {
	aoc.RunSolution(part1, "\n\n")
	aoc.RunExample(part2, "\n\n")
}

func part2(input []string) (int, error) {
	return solution(input, 10_000, false)
}

func part1(input []string) (int, error) {
	return solution(input, 20, true)
}

func solution(input []string, rounds int, getsBored bool) (int, error) {
	h, err := newHerdFromInput(input)
	if err != nil {
		return 0, err
	}

	for i := 0; i < rounds; i++ {
		if err := h.round(getsBored); err != nil {
			return 0, err
		}

		if !_roundSummary {
			continue
		}

		if getsBored { // part 1
			fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", i+1)
			for j, m := range h {
				fmt.Printf("Monkey %d: %v\n", j, m.items)
			}
		}
		if !getsBored && (i == 0 || i == 19 || (i+1)%1000 == 0) { // part 2
			fmt.Printf("== After round %d ==\n", i+1)
			for j, m := range h {
				fmt.Printf("Monkey %d inspected items %d times.\n", j, m.inspections)
			}
		}
	}

	result := []int{}
	for _, m := range h {
		result = append(result, m.inspections)
	}
	sort.Ints(result)
	fmt.Println(result)
	return result[len(result)-1] * result[len(result)-2], nil
}

type herd []*monkey

func (h herd) round(getsBored bool) error {
	for i, m := range h {
		msg("Monkey %d:\n", i)
		for _, item := range m.items {
			msg("  Monkey inspects an item with a worry level of  %d.\n", item)
			if err := m.inspect(); err != nil {
				return err
			}
			msg("    Worry level goes from %d to %d.\n", item, m.items[0])

			if getsBored {
				m.getBored()
				msg("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", m.items[0])
			}

			target := m.test(m.items[0])
			msg("    Item with worry level %d is thrown to monkey %d.\n", m.items[0], target)
			m.throw(h[target])
		}
	}

	return nil
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
	inspections int
	operation   func(int) (int, error)
	test        func(int) int
	items       []int
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

func (m *monkey) inspect() error {
	m.inspections++
	item, err := m.operation(m.items[0])
	if err != nil {
		return err
	}
	m.items[0] = item
	return nil
}

func (m *monkey) getBored() {
	m.items[0] /= 3
}

func (m *monkey) throw(other *monkey) {
	item := m.items[0]
	m.items = m.items[1:]
	other.items = append(other.items, item)
}

func msg(format string, v ...any) {
	if _verbose {
		fmt.Printf(format, v...)
	}
}
