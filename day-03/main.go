package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

func main() {
	// Part 1
	println(sumItemPriorities())
}

func sumItemPriorities() (int, error) {
	lines, err := input()
	if err != nil {
		return 0, err
	}

	result := 0
	for _, line := range lines {
		r, err := getTheOnlyDuplicate(line)
		if err != nil {
			return 0, err
		}

		p, err := priority(r)
		if err != nil {
			return 0, err
		}

		result += p
	}

	return result, nil
}

func priority(r rune) (int, error) {
	switch {
	case 'A' <= r && r <= 'Z':
		// 65 <= int(r) = 90
		return int(r) - 38, nil
	case 'a' <= r && r <= 'z':
		// 97 <= int(r)<= 122
		return int(r) - 96, nil
	}
	return 0, errors.New("illegal rune")
}

func getTheOnlyDuplicate(line string) (rune, error) {
	m := map[rune]struct{}{}

	for _, r := range line[:len(line)/2] {
		m[r] = struct{}{}
	}

	for _, r := range line[len(line)/2:] {
		if _, ok := m[r]; ok {
			return r, nil
		}
	}

	return 0, errors.New("not found")
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
