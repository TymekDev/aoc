package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

func main() {
	// Part 1
	println(totalScore(input()))
}

func totalScore(lines []string) int {
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

func parseLine(line string) (them Shape, me Shape) {
	l := strings.Split(line, " ")
	return m[l[0]], m[l[1]]
}

func roundScore(them, me Shape) int {
	return outcome(them, me) + int(me)
}

func outcome(them, me Shape) int {
	switch {
	case me == Rock && them == Scissors,
		me == Scissors && them == Paper,
		me == Paper && them == Rock:
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
