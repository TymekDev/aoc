package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStacks(t *testing.T) {
	s := newStacks([]string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	})

	require.Equal(t, 3, len(s))

	assert.Equal(t, "N", s[0].label)
	assert.Equal(t, "Z", s[0].below.label)
	assert.Equal(t, (*crate)(nil), s[0].below.below)

	assert.Equal(t, "D", s[1].label)
	assert.Equal(t, "C", s[1].below.label)
	assert.Equal(t, "M", s[1].below.below.label)
	assert.Equal(t, (*crate)(nil), s[1].below.below.below)

	assert.Equal(t, "P", s[2].label)
	assert.Equal(t, (*crate)(nil), s[2].below)
}

func TestMove(t *testing.T) {
	a := &crate{"A", nil}
	b := &crate{"B", nil}
	c := &crate{"C", b}

	s := stacks{a, c, nil} // => a, cb, _

	s.move(0, 1) // => _, acb _

	assert.Equal(t, s[0], (*crate)(nil))
	assert.Equal(t, s[1], a)

	s.move(1, 0) // => a, cb, _
	assert.Equal(t, s[0], a)
	assert.Equal(t, s[1], c)

	s.move(1, 0) // => ca, b, _
	assert.Equal(t, s[0], c)
	assert.Equal(t, s[1], b)

	s.move(0, 2) // => bca, _, _
	assert.Equal(t, s[0], a)
	assert.Equal(t, s[1], b)
	assert.Equal(t, s[2], c)
}
