package main

import (
	"fmt"

	"git.sr.ht/~tymek/aoc-2022"
)

func main() {
	aoc.RunExample(solution, "\n")
	aoc.RunSolution(solution, "\n")
}

func solution(input []string) (string, error) {
	nVisible := 0
	maxScore := 0
	for row, line := range input {
		for col := range line {
			if row == 0 || row == len(input)-1 || col == 0 || col == len(line)-1 {
				nVisible++
				continue
			}

			height := h(line[col])
			visible := false
			score := 1

			// horizontal
			for i := col - 1; i >= 0; i-- {
				if h(line[i]) >= height {
					score *= col - i
					break
				}
				if i == 0 {
					score *= col - i
					visible = true
				}
			}

			for i := col + 1; i < len(line); i++ {
				if h(line[i]) >= height {
					score *= i - col
					break
				}
				if i == len(line)-1 {
					score *= i - col
					visible = true
				}
			}

			// horizontal
			for i := row - 1; i >= 0; i-- {
				if h(input[i][col]) >= height {
					score *= row - i
					break
				}
				if i == 0 {
					score *= row - i
					visible = true
				}
			}

			for i := row + 1; i < len(input); i++ {
				if h(input[i][col]) >= height {
					score *= i - row
					break
				}
				if i == len(input)-1 {
					score *= i - row
					visible = true
				}
			}

			if visible {
				nVisible++
			}

			if score > maxScore {
				maxScore = score
			}
		}
	}

	return fmt.Sprintf("%d %d\n", nVisible, maxScore), nil
}

func h(b byte) int {
	return int(b) - 48
}
