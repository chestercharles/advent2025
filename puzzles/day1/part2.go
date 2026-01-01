package day1

import (
	"strconv"
	"strings"

	"github.com/chestercharles/advent2025/internal/dial"
)




func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	myDial := dial.NewDial()

	zeroes := 0
	for _, element := range lines {
		dir, n, err := parseCode(element)

		if (err != nil) {
			panic(err)
		}

		if (dir == "L") {
			zeroes = zeroes + myDial.TurnLeft(n)
		}

		if (dir == "R") {
			zeroes = zeroes + myDial.TurnRight(n)
		}
	}
	
	return strconv.Itoa(zeroes), nil
}