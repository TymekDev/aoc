package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalScore(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	assert.Equal(t, 15, totalScore(lines))
}

func TestRoundScore(t *testing.T) {
	tests := []struct {
		name   string
		me     Shape
		them   Shape
		result int
	}{
		{"Lose Rock", Rock, Paper, 1},
		{"Lose Paper", Paper, Scissors, 2},
		{"Lose Scissors", Scissors, Rock, 3},
		{"Win Rock", Rock, Scissors, 7},
		{"Win Paper", Paper, Rock, 8},
		{"Win Scissors", Scissors, Paper, 9},
		{"Draw Rock", Rock, Rock, 4},
		{"Draw Paper", Paper, Paper, 5},
		{"Draw Scissors", Scissors, Scissors, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, roundScore(tt.them, tt.me))
		})
	}
}

func TestOutcome(t *testing.T) {
	tests := []struct {
		name   string
		me     Shape
		them   Shape
		result int
	}{
		{"Scissors < Rock", Scissors, Rock, 0},
		{"Paper < Scissors", Paper, Scissors, 0},
		{"Rock < Paper", Rock, Paper, 0},
		{"Paper == Paper", Paper, Paper, 3},
		{"Scissors == Scissors", Scissors, Scissors, 3},
		{"Rock == Rock", Rock, Rock, 3},
		{"Rock > Scissors", Rock, Scissors, 6},
		{"Scissors > Paper", Scissors, Paper, 6},
		{"Paper > Rock", Paper, Rock, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, outcome(tt.them, tt.me))
		})
	}
}
