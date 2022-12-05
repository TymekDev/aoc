package aoc

import (
	"io/ioutil"
	"log"
	"strings"
)

type solution[T answer] func([]string) (T, error)

type answer interface {
	int | string
}

func RunSolution[T answer](f solution[T], sep string) {
	run(f, input, sep)
}

func RunExample[T answer](f solution[T], sep string) {
	run(f, example, sep)
}

func run[T answer](f solution[T], i func(string) ([]string, error), sep string) {
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
