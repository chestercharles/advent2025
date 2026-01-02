package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	sum := 0

	for _, rangeStr := range strings.Split(input, ",") {
		rangeStartEnd := strings.Split(rangeStr, "-")
		start := rangeStartEnd[0]
		end := rangeStartEnd[1]

		r, err := makeArr(start, end)

		if err != nil {
			return "", err
		}

		for _, n := range r {
			if silly(n) {
				sum += n
			}
		}

	}

	return strconv.Itoa(sum), nil
}

func silly(id int) bool {
	str := fmt.Sprint(id)
	l := len(str)
	if l%2 == 1 {
		return false
	}

	mid := l / 2

	s := str[:mid]
	e := str[mid:]

	return s == e
}
