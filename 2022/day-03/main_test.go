package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPriorityForTheOnlyCommonRune(t *testing.T) {
	tests := []struct {
		lines  [3]string
		result rune
	}{
		{
			[3]string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			'r',
		},
		{
			[3]string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			'Z',
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.result), func(t *testing.T) {
			p, err := priority(tt.result)
			require.NoError(t, err)
			result, err := getPriorityForTheOnlyCommonRune(tt.lines[0], tt.lines[1], tt.lines[2])
			require.NoError(t, err)
			assert.Equal(t, p, result)
		})
	}
}

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
