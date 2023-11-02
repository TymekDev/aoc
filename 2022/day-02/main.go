package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Shape int

// NOTE: x wins with y if and only if (4-x)%3 == 3-y
const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

func main() {
	// Part 1
	println(totalScore(input(), parseLinePart1))

	// Part 2
	println(totalScore(input(), parseLinePart2))
}

func totalScore(lines []string, parseLine lineParser) int {
	result := 0
	for _, line := range lines {
		them, me := parseLine(line)
		result += roundScore(them, me)
	}
	return result
}

var m map[string]Shape = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

type lineParser func(line string) (them Shape, me Shape)

func parseLinePart1(line string) (them Shape, me Shape) {
	l := strings.Split(line, " ")
	return m[l[0]], m[l[1]]
}

func parseLinePart2(line string) (them Shape, me Shape) {
	l := strings.Split(line, " ")
	them = m[l[0]]
	switch l[1] {
	case "X": // lose
		me = 3 - (4-them)%3
	case "Y": // draw
		me = them
	case "Z": // win
		me = them%3 + 1
	}

	return them, me
}

func roundScore(them, me Shape) int {
	return outcome(them, me) + int(me)
}

func outcome(them, me Shape) int {
	switch {
	case (4-me)%3 == 3-them:
		return 6
	case me == them:
		return 3
	}
	return 0
}

func input() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")
}
