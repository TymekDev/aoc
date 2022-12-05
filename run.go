package aoc

import (
	"io/ioutil"
	"log"
	"strings"
)

type solution func([]string) (int, error)

func RunSolution(f solution, sep string) {
	run(f, input, sep)
}

func RunExample(f solution, sep string) {
	run(f, example, sep)
}

func run(f solution, i func(string) ([]string, error), sep string) {
	input, err := i(sep)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := f(input)
	if err != nil {
		log.Fatalln(err)
	}

	println(result)
}

func input(sep string) ([]string, error) {
	return read("input.txt", sep)
}

func example(sep string) ([]string, error) {
	return read("input_example.txt", sep)
}

func read(fileName, sep string) ([]string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSuffix(string(b), "\n"), sep), nil
}
