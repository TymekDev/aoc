package main

import (
	"math"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	aoc.RunExample(part1, "\n")
}

func part1(input []string) (int, error) {
	h, start := newHillFromInput(input)
	h.traverse(start, map[*point]struct{}{})
	return h.best, nil
}

type hill struct {
	best      int
	elevation [][]*point
}

type point struct {
	value rune
	x     int
	y     int
}

func newHillFromInput(input []string) (*hill, *point) {
	var start *point
	h := &hill{best: math.MaxInt, elevation: make([][]*point, len(input))}
	for i, line := range input {
		h.elevation[i] = make([]*point, len(line))
		for j, r := range line {
			h.elevation[i][j] = &point{r, i, j}
			switch r {
			case 'S':
				start = h.elevation[i][j]
				h.elevation[i][j].value = 'a' - 1
			case 'E':
				h.elevation[i][j].value = 'z' + 1
			}
		}
	}

	return h, start
}

func (h *hill) neighbors(p *point) []*point {
	result := []*point{}
	if next, ok := h.get(p.x-1, p.y); ok {
		result = append(result, next)
	}
	if next, ok := h.get(p.x+1, p.y); ok {
		result = append(result, next)
	}
	if next, ok := h.get(p.x, p.y-1); ok {
		result = append(result, next)
	}
	if next, ok := h.get(p.x, p.y+1); ok {
		result = append(result, next)
	}
	return result
}

func (h *hill) get(i, j int) (*point, bool) {
	if i < 0 || j < 0 || i >= len(h.elevation) || j >= len(h.elevation[0]) {
		return nil, false
	}
	return h.elevation[i][j], true
}

func (h *hill) traverse(p *point, visited map[*point]struct{}) {
	if len(visited) > h.best {
		return
	}
	visited[p] = struct{}{}

	if p.value == 'z'+1 {
		if steps := len(visited) - 1; steps < h.best {
			h.best = steps
		}
		return
	}

	for _, next := range h.neighbors(p) {
		if _, ok := visited[next]; !ok && (next.value == p.value || next.value == p.value+1) {
			h.traverse(next, copyMap(visited))
		}
	}
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
