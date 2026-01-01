package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chestercharles/advent2025/internal/dial1"
)


func parseCode(s string) (string, int, error) {
	if (len(s) < 2) {
		return "", 0, fmt.Errorf("too short")
	}

	dir := s[:1]
	n, err := strconv.Atoi(s[1:])
	if (err != nil) {
		return dir, 0, err
	}

	return dir, n, err
}



func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	myDial := dial1.NewDial()

	zeroes := 0
	for _, element := range lines {
		dir, n, err := parseCode(element)

		if (err != nil) {
			panic(err)
		}

		if (dir == "L") {
			myDial.TurnLeft(n)
		}

		if (dir == "R") {
			myDial.TurnRight(n)
		}

		if (myDial.Read() == 0) {
			zeroes++
		}
	}
	
	return strconv.Itoa(zeroes), nil
}