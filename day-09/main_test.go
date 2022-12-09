package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointVisited(t *testing.T) {
	tests := []struct {
		// point starts at 0,0 and result is incremented to account for that
		moves  [][2]int
		result int
	}{
		{
			[][2]int{{0, 1}, {0, 2}},
			2,
		},
		{
			[][2]int{{0, 0}, {0, 0}},
			0,
		},
		{
			[][2]int{{1, 0}, {0, 0}, {1, 0}},
			2,
		},
		{
			[][2]int{{1, 0}, {1, 2}, {-2, -3}, {-3, -2}},
			4,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := newPoint(0, 0)
			for _, m := range tt.moves {
				p.move(m[0], m[1])
			}
			assert.Equal(t, tt.result+1, p.uniqueVisitedCount())
		})
	}
}

func TestPointFollow(t *testing.T) {
	tests := []struct {
		p      *point
		target *point
		result *point
	}{
		// Don't move if we are within 1 field horizontally, vertically or diagonally
		{newPoint(0, 0), newPoint(-1, -1), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(-1, 0), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(-1, 1), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(0, -1), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(0, 0), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(0, 1), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(1, -1), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(1, 0), newPoint(0, 0)},
		{newPoint(0, 0), newPoint(1, 1), newPoint(0, 0)},
		// vertical movement
		{newPoint(0, 0), newPoint(0, 2), newPoint(0, 1)},
		{newPoint(0, 20), newPoint(0, 18), newPoint(0, 19)},
		// horizontal movement
		{newPoint(3, 0), newPoint(3, -2), newPoint(3, -1)},
		{newPoint(3, 0), newPoint(3, 2), newPoint(3, 1)},
		// diagonal movement
		{newPoint(0, 0), newPoint(-2, -1), newPoint(-1, -1)},
		{newPoint(0, 0), newPoint(-2, 1), newPoint(-1, 1)},
		{newPoint(0, 0), newPoint(-1, -2), newPoint(-1, -1)},
		{newPoint(0, 0), newPoint(1, 2), newPoint(1, 1)},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.p.follow(tt.target)
			assert.Equal(t, tt.result.x, tt.p.x)
			assert.Equal(t, tt.result.y, tt.p.y)
		})
	}
}
