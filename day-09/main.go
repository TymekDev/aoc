package main

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
	p.x, p.y = x, y
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
