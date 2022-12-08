package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestH(t *testing.T) {
	tests := []struct {
		b      byte
		result int
	}{
		{byte('0'), 0},
		{byte('1'), 1},
		{byte('2'), 2},
		{byte('3'), 3},
		{byte('4'), 4},
		{byte('5'), 5},
		{byte('6'), 6},
		{byte('7'), 7},
		{byte('8'), 8},
		{byte('9'), 9},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.result, h(tt.b))
		})
	}
}
