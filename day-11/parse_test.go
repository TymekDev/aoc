package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseItems(t *testing.T) {
	items, err := parseItems("  Starting items: 79, 98")
	require.NoError(t, err)
	assert.ElementsMatch(t, []int{79, 98}, items)
}

func TestParseOperation(t *testing.T) {
	operation, err := parseOperation("  Operation: new = old * 19")
	require.NoError(t, err)
	result, err := operation(5)
	require.NoError(t, err)
	assert.Equal(t, 19*5, result)
}

func TestParseTest(t *testing.T) {
	test, err := parseTest(
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
	)
	require.NoError(t, err)
	assert.Equal(t, 2, test(23))
	assert.Equal(t, 3, test(24))
}
