package main

import (
	"strconv"
	"strings"

	"github.com/TymekDev/aoc/2022"
)

func main() {
	aoc.RunSolution(part1, "\n")
	aoc.RunSolution(part2, "\n")
}

func part2(input []string) (string, error) {
	c, err := solution(input)
	if err != nil {
		return "", err
	}

	return c.render(), nil
}

func part1(input []string) (int, error) {
	c, err := solution(input)
	if err != nil {
		return 0, err
	}

	return c.result, nil
}

func solution(input []string) (*cycler, error) {
	c := newCycler()
	for _, line := range input {
		_, value, found := strings.Cut(line, " ")
		if found {
			x, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			c.add(x)
		} else {
			c.tick()
		}
	}

	return c, nil
}

type cycler struct {
	cycle    int
	register int
	result   int
	targets  []int
	screen   []string
}

func newCycler() *cycler {
	return &cycler{
		register: 1,
		targets:  []int{20, 60, 100, 140, 180, 220},
		screen:   []string{},
	}
}

func (c *cycler) add(i int) {
	c.tick()
	c.tick()
	c.register += i
}

func (c *cycler) tick() {
	if c.register-1 <= c.cycle%40 && c.cycle%40 <= c.register+1 {
		c.screen = append(c.screen, "#")
	} else {
		c.screen = append(c.screen, ".")
	}

	c.cycle++

	if len(c.targets) == 0 {
		return
	}

	if target := c.targets[0]; c.cycle >= target {
		c.result += target * c.register
		c.targets = c.targets[1:]
	}
}

func (c *cycler) render() string {
	i := 0
	var sb strings.Builder
	for 40*(i+1) <= len(c.screen) {
		sb.WriteString(strings.Join(c.screen[40*i:40*(i+1)], ""))
		sb.WriteString("\n")
		i++
	}
	return sb.String()
}
