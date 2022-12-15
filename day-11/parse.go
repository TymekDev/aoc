package main

import (
	"errors"
	"strconv"
	"strings"
)

// line: "  Starting items: 79, 98"
func parseItems(line string) ([]int, error) {
	items := strings.Fields(strings.ReplaceAll(line, ",", ""))[2:]
	result := make([]int, len(items))
	for i, item := range items {
		value, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}

		result[i] = value
	}

	return result, nil
}

// line: "  Operation: new = old * 19"
func parseOperation(line string) (func(int) (int, error), error) {
	fields := strings.Fields(line)
	lhs := fields[3]
	op := fields[4]
	rhs := fields[5]

	switch {
	case lhs == "old" && rhs == "old":
		return func(x int) (int, error) {
			switch op {
			case "+":
				return x + x, nil
			case "-":
				return x - x, nil
			case "*":
				return x * x, nil
			case "/":
				return x / x, nil
			}
			return 0, errors.New("operation not recognized: " + op)
		}, nil
	case lhs == "old" && rhs != "old":
		v, err := strconv.Atoi(rhs)
		if err != nil {
			return nil, err
		}
		return func(x int) (int, error) {
			switch op {
			case "+":
				return x + v, nil
			case "-":
				return x - v, nil
			case "*":
				return x * v, nil
			case "/":
				return x / v, nil
			}
			return 0, errors.New("operation not recognized: " + op)
		}, nil
	}

	return nil, errors.New("unexpected case")
}

// test: "  Test: divisible by 23"
// t: "    If true: throw to monkey 2"
// f: "    If false: throw to monkey 3"
func parseTest(test, t, f string) (func(int) int, error) {
	testFields := strings.Fields(test)
	testValue, err := strconv.Atoi(testFields[len(testFields)-1])
	if err != nil {
		return nil, err
	}

	tFields := strings.Fields(t)
	tValue, err := strconv.Atoi(tFields[len(tFields)-1])
	if err != nil {
		return nil, err
	}

	fFields := strings.Fields(f)
	fValue, err := strconv.Atoi(fFields[len(fFields)-1])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		if x%testValue == 0 {
			return tValue
		}
		return fValue
	}, nil
}
