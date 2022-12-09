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
			1,
		},
		{
			[][2]int{{1, 0}, {0, 0}, {1, 0}},
			1,
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
