package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	tests := []struct {
		input  string
		result int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			s := newScanner(tt.input)
			result, err := s.scan()
			require.NoError(t, err)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestIsMarker(t *testing.T) {
	tests := []struct {
		input          string
		markerPosition int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			s := newScanner(tt.input)
			s.position = tt.markerPosition
			assert.Equal(t, true, s.isMarker())
		})
	}
}
