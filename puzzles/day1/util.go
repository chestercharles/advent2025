package day1

import (
	"fmt"
	"strconv"
)

func parseCode(s string) (string, int, error) {
	if len(s) < 2 {
		return "", 0, fmt.Errorf("too short")
	}

	dir := s[:1]
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return dir, 0, err
	}

	return dir, n, err
}
