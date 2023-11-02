package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	println(top1())

	// Part 2
	println(top3())
}

func top3() int {
	calories := calcCalories()
	sort.Ints(calories)

	return calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
}

func calcCalories() []int {
	currentElf := 0
	calories := []int{0}
	for _, line := range strings.Split(input(), "\n") {
		if line == "" {
			currentElf++
			calories = append(calories, 0)
			continue
		}

		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}

		calories[currentElf] += x
	}

	return calories
}

func top1() int {
	currentElf := 0
	maxElf := 0
	calories := []int{0}
	for _, line := range strings.Split(input(), "\n") {
		if line == "" {
			if calories[currentElf] > calories[maxElf] {
				maxElf = currentElf
			}
			currentElf++
			calories = append(calories, 0)
			continue
		}

		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}

		calories[currentElf] += x
	}

	return calories[maxElf]
}

func input() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}
