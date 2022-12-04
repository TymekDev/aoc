package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOverlapFull(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{"1-4,2-3", true},
		{"2-3,1-4", true},
		{"1-1,1-1", true},
		{"1-2,1-2", true},
		{"6-6,4-6", true},
		{"4-6,6-6", true},
		{"1-3,2-4", false},
		{"5-7,2-4", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			rp, err := newRangePairFromInput(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.result, rp.overlapFull())
		})
	}
}

func TestOverlapAny(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{"1-4,2-3", true},
		{"2-3,1-4", true},
		{"1-1,1-1", true},
		{"1-2,1-2", true},
		{"6-6,4-6", true},
		{"4-6,6-6", true},
		{"1-3,2-4", true},
		{"5-7,2-4", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			rp, err := newRangePairFromInput(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.result, rp.overlapAny())
		})
	}
}

func TestContainsFully(t *testing.T) {
	tests := []struct {
		outer  string
		inner  string
		result bool
	}{
		{"1-4", "2-3", true},
		{"1-1", "1-1", true},
		{"1-2", "1-2", true},
		{"1-3", "2-2", true},
		{"1-3", "2-3", true},
		{"1-3", "2-4", false},
		{"2-2", "1-4", false},
	}

	for _, tt := range tests {
		t.Run(tt.outer+" fully contains "+tt.inner, func(t *testing.T) {
			outer, err := newRangeFromInput(tt.outer)
			require.NoError(t, err)
			inner, err := newRangeFromInput(tt.inner)
			require.NoError(t, err)

			assert.Equal(t, tt.result, outer.containsFully(inner))
		})
	}
}

func TestContainsEnd(t *testing.T) {
	tests := []struct {
		r      string
		other  string
		result bool
	}{
		{"1-4", "2-3", true},
		{"1-1", "1-1", true},
		{"1-2", "1-2", true},
		{"1-3", "2-2", true},
		{"1-3", "2-3", true},
		{"1-3", "2-4", true},
		{"2-2", "1-4", false},
	}

	for _, tt := range tests {
		t.Run(tt.r+" contains end of "+tt.other, func(t *testing.T) {
			r, err := newRangeFromInput(tt.r)
			require.NoError(t, err)
			other, err := newRangeFromInput(tt.other)
			require.NoError(t, err)

			assert.Equal(t, tt.result, r.containsEnd(other))
		})
	}
}
