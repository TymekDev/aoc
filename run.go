package aoc

import (
	"io/ioutil"
	"log"
	"strings"
)

type solution func([]string) (int, error)

func RunSolution(f solution) {
	run(f, input)
}

func RunExample(f solution) {
	run(f, example)
}

func run(f solution, i func() ([]string, error)) {
	input, err := i()
	if err != nil {
		log.Fatalln(err)
	}

	result, err := f(input)
	if err != nil {
		log.Fatalln(err)
	}

	println(result)
}

func input() ([]string, error) {
	return read("input.txt")
}

func example() ([]string, error) {
	return read("input_example.txt")
}

func read(fileName string) ([]string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSuffix(string(b), "\n"), "\n"), nil
}
