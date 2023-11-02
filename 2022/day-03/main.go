package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

func main() {
	// Part 1
	println(sumItemPriorities())

	// Part 2
	println(sumBadgesPriorities())
}

func sumBadgesPriorities() (int, error) {
	lines, err := input()
	if err != nil {
		return 0, err
	}

	result := 0
	for i := 0; i < len(lines); i += 3 {
		p, err := getPriorityForTheOnlyCommonRune(lines[i], lines[i+1], lines[i+2])
		if err != nil {
			return 0, err
		}

		result += p
	}

	return result, nil
}

func getPriorityForTheOnlyCommonRune(line1, line2, line3 string) (int, error) {
	occurences := [52][3]int{}

	for i, line := range []string{line1, line2, line3} {
		for _, r := range line {
			p, err := priority(r)
			if err != nil {
				return 0, err
			}

			occurences[p-1][i]++
		}
	}

	for i, occ := range occurences {
		if occ[0] > 0 && occ[1] > 0 && occ[2] > 0 {
			return i + 1, nil
		}
	}

	return 0, errors.New("not found")
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
