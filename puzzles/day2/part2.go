package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(input string) (string, error) {
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
			if silly2(n) {
				sum += n
			}
		}

	}

	return strconv.Itoa(sum), nil
}

func silly2(id int) bool {
	str := fmt.Sprint(id)
	l := len(str)
	if l < 1 {
		return false
	}

	doubled := str + str
	sub := doubled[1 : len(doubled)-1]

	return strings.Contains(sub, str)
}
