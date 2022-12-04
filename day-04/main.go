package main

import (
	"errors"
	"strconv"
	"strings"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	// Part 1
	aoc.RunSolution(countFullOverlaps)
}

func countFullOverlaps(input []string) (int, error) {
	result := 0
	for _, line := range input {
		rp, err := newRangePairFromInput(line)
		if err != nil {
			return 0, err
		}

		if rp.overlapFull() {
			result++
		}
	}

	return result, nil
}

type rangePair struct {
	r1 *Range
	r2 *Range
}

func newRangePairFromInput(input string /* <r1.lower>-<r1.upper>,<r2.lower>-<r2.upper> */) (*rangePair, error) {
	rp := strings.Split(input, ",")
	if len(rp) != 2 {
		return nil, errors.New("something went wrong")
	}

	r1, err := newRangeFromInput(rp[0])
	if err != nil {
		return nil, err
	}

	r2, err := newRangeFromInput(rp[1])
	if err != nil {
		return nil, err
	}

	return &rangePair{r1: r1, r2: r2}, nil
}

func (rp *rangePair) overlapFull() bool {
	return rp.r1.containsFully(rp.r2) || rp.r2.containsFully(rp.r1)
}

type Range struct {
	Lower int
	Upper int
}

func newRangeFromInput(input string /* <lower>-<upper> */) (*Range, error) {
	r := strings.Split(input, "-")
	if len(r) != 2 {
		return nil, errors.New("something went wrong")
	}

	lower, err := strconv.Atoi(r[0])
	if err != nil {
		return nil, err
	}

	upper, err := strconv.Atoi(r[1])
	if err != nil {
		return nil, err
	}

	return &Range{Lower: lower, Upper: upper}, nil
}

func (r *Range) containsFully(inner *Range) bool {
	return r.Lower <= inner.Lower && inner.Upper <= r.Upper
}
