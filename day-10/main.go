package main

import (
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	aoc.RunSolution(part1, "\n")
}

func part1(input []string) (int, error) {
	c := newCycler()
	for _, line := range input {
		_, value, found := strings.Cut(line, " ")
		if found {
			x, err := strconv.Atoi(value)
			if err != nil {
				return 0, err
			}

			c.add(x)
		} else {
			c.tick()
		}
	}

	return c.result, nil
}

type cycler struct {
	cycle    int
	register int
	result   int
	targets  []int
}

func newCycler() *cycler {
	return &cycler{register: 1, targets: []int{20, 60, 100, 140, 180, 220}}
}

func (c *cycler) add(i int) {
	c.tick()
	c.tick()
	c.register += i
}

func (c *cycler) tick() {
	c.cycle++

	if len(c.targets) == 0 {
		return
	}

	if target := c.targets[0]; c.cycle >= target {
		c.result += target * c.register
		c.targets = c.targets[1:]
	}
}
