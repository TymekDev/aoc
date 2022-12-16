package main

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// line: "  Starting items: 79, 98"
func parseItems(line string) ([]*big.Int, error) {
	items := strings.Fields(strings.ReplaceAll(line, ",", ""))[2:]
	result := make([]*big.Int, len(items))
	for i, item := range items {
		value, ok := (&big.Int{}).SetString(item, 10)
		if !ok {
			return nil, errors.New("couldn't do the thing")
		}

		result[i] = value
	}

	return result, nil
}

// line: "  Operation: new = old * 19"
func parseOperation(line string) (func(*big.Int) (*big.Int, error), error) {
	fields := strings.Fields(line)
	lhs := fields[3]
	op := fields[4]
	rhs := fields[5]

	switch {
	case lhs == "old" && rhs == "old":
		return func(x *big.Int) (*big.Int, error) {
			result := &big.Int{}
			switch op {
			case "+":
				result.Add(x, x)
			case "-":
			case "*":
				result.Mul(x, x)
			case "/":
				result.SetInt64(1)
			default:
				return nil, errors.New("operation not recognized: " + op)
			}
			return result, nil
		}, nil
	case lhs == "old" && rhs != "old":
		v, ok := (&big.Int{}).SetString(rhs, 10)
		if !ok {
			return nil, errors.New("couldn't do the thing")
		}
		return func(x *big.Int) (*big.Int, error) {
			result := &big.Int{}
			switch op {
			case "+":
				result.Add(x, v)
			case "-":
				result.Sub(x, v)
			case "*":
				result.Mul(x, v)
			case "/":
				result.Div(x, v)
			default:
				return nil, errors.New("operation not recognized: " + op)
			}
			return result, nil
		}, nil
	}

	return nil, errors.New("unexpected case")
}

// test: "  Test: divisible by 23"
// t: "    If true: throw to monkey 2"
// f: "    If false: throw to monkey 3"
func parseTest(test, t, f string) (func(*big.Int) int, error) {
	testFields := strings.Fields(test)
	testValue, ok := (&big.Int{}).SetString(testFields[len(testFields)-1], 10)
	if !ok {
		return nil, errors.New("couldn't do the thing")
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

	return func(x *big.Int) int {
		if (&big.Int{}).Mod(x, testValue).Cmp(&big.Int{}) == 0 {
			return tValue
		}
		return fValue
	}, nil
}
