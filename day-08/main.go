package main

import (
	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	aoc.RunSolution(part1, "\n")
}

func part1(input []string) (int, error) {
	result := 0
	for row, line := range input {
		for col := range line {
			height := h(line[col])
			if row == 0 || row == len(input)-1 || col == 0 || col == len(line)-1 {
				goto NEXT
			}

			// horizontal
			for i := col - 1; i >= 0; i-- {
				if h(line[i]) >= height {
					break
				}
				if i == 0 {
					goto NEXT
				}
			}

			for i := col + 1; i < len(line); i++ {
				if h(line[i]) >= height {
					break
				}
				if i == len(line)-1 {
					goto NEXT
				}
			}

			// horizontal
			for i := row - 1; i >= 0; i-- {
				if h(input[i][col]) >= height {
					break
				}
				if i == 0 {
					goto NEXT
				}
			}

			for i := row + 1; i < len(input); i++ {
				if h(input[i][col]) >= height {
					break
				}
				if i == len(input)-1 {
					goto NEXT
				}
			}

			continue

		NEXT:
			result++
		}
	}

	return result, nil
}

func h(b byte) int {
	return int(b) - 48
}
