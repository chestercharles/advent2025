package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part2(input string) (string, error) {

	banks := strings.Split(input, "\n")

	sum := 0
	for _, bank := range banks {
		j, err := joltage2(bank)
		if err != nil {
			return "", err
		}
		sum = sum + j
	}
	return fmt.Sprint(sum), nil
}

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() int {
	i := len(s.items) - 1

	item := s.items[i]

	s.items = s.items[:i]

	return item
}

func joltage2(bank string) (int, error) {
	l := len(bank)

	arr, err := makeArr(bank)

	if err != nil {
		return 0, err
	}

	s := &Stack{items: arr}
	d := arr[0]
	remove := len(arr) - 12
	for !s.IsEmpty() && s.Peek() < d && remove > 0 {

	}

	max := slices.Max(arr)

	i := slices.Index(arr, max)

	if i == l-1 {
		nextMax := slices.Max(arr[:l-1])
		return strconv.Atoi(fmt.Sprint(nextMax) + fmt.Sprint(max))
	}

	nextMax := slices.Max(arr[i+1:])

	return strconv.Atoi(fmt.Sprint(max) + fmt.Sprint(nextMax))
}
