package main

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	aoc.RunSolution(part1, "\n")
	aoc.RunSolution(part2, "\n")
}

func part2(input []string) (int, error) {
	r := newRope()
	for _, line := range input {
		x, y, n, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		for i := 0; i < n; i++ {
			r.move(x, y)
		}
	}

	return r[len(r)-1].uniqueVisitedCount(), nil
}

func part1(input []string) (int, error) {
	head, tail := newPoint(0, 0), newPoint(0, 0)
	for _, line := range input {
		x, y, n, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		for i := 0; i < n; i++ {
			head.move(x, y)
			tail.follow(head)
		}
	}

	return tail.uniqueVisitedCount(), nil
}

func parseLine(line string) (x int, y int, times int, err error) {
	f := strings.Fields(line)
	if len(f) != 2 {
		return 0, 0, 0, errors.New("incorrect line")
	}

	n, err := strconv.Atoi(f[1])
	if err != nil {
		return 0, 0, 0, err
	}

	switch f[0] {
	case "R":
		return 1, 0, n, nil
	case "L":
		return -1, 0, n, nil
	case "U":
		return 0, 1, n, nil
	case "D":
		return 0, -1, n, nil
	}

	return 0, 0, 0, errors.New("incorrect line")
}

type rope [10]*point

func newRope() *rope {
	r := &rope{}
	for i := 0; i < len(r); i++ {
		r[i] = newPoint(0, 0)
	}
	return r
}

func (r *rope) move(x, y int) {
	r[0].move(x, y)
	for i := 1; i < len(r); i++ {
		r[i].follow(r[i-1])
	}
}

type point struct {
	x       int
	y       int
	visited map[int]map[int]struct{}
}

func newPoint(x, y int) *point {
	p := &point{visited: map[int]map[int]struct{}{}}
	p.move(x, y)
	return p
}

func (p *point) move(x, y int) {
	p.x, p.y = p.x+x, p.y+y
	if _, ok := p.visited[p.x]; !ok {
		p.visited[p.x] = map[int]struct{}{}
	}
	p.visited[p.x][p.y] = struct{}{}
}

func (p *point) uniqueVisitedCount() int {
	result := 0
	for _, v := range p.visited {
		result += len(v)
	}
	return result
}

func (p *point) follow(target *point) {
	if math.Abs(float64(target.x-p.x)) <= 1 && math.Abs(float64(target.y-p.y)) <= 1 {
		return
	}

	p.move(compare(target.x, p.x), compare(target.y, p.y))
}

func compare(x, y int) int {
	if x > y {
		return 1
	}
	if x == y {
		return 0
	}
	return -1
}
