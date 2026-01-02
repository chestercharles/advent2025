package day1

import (
	"strconv"
	"strings"

	"github.com/chestercharles/advent2025/internal/dial"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	myDial := dial.NewDial()

	zeroes := 0
	for _, element := range lines {
		dir, n, err := parseCode(element)

		if err != nil {
			panic(err)
		}

		if dir == "L" {
			myDial.TurnLeft(n)
		}

		if dir == "R" {
			myDial.TurnRight(n)
		}

		if myDial.Read() == 0 {
			zeroes++
		}
	}

	return strconv.Itoa(zeroes), nil
}
