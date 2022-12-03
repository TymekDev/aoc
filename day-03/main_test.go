package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPriority(t *testing.T) {
	tests := []struct {
		r      rune
		result int
	}{
		{'a', 1},
		{'b', 2},
		{'z', 26},
		{'A', 27},
		{'B', 28},
		{'Z', 52},
	}

	for _, tt := range tests {
		t.Run(string(tt.r), func(t *testing.T) {
			result, err := priority(tt.r)
			require.NoError(t, err)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestGetTheOnlyDuplicate(t *testing.T) {
	tests := []struct {
		line   string
		result rune
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
		{"PmmdzqPrVvPwwTWBwg", 'P'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
		{"ttgJtRGJQctTZtZT", 't'},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result, err := getTheOnlyDuplicate(tt.line)
			require.NoError(t, err)
			assert.Equal(t, tt.result, result)
		})
	}
}
