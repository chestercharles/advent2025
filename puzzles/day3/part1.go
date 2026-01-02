package day3

import (
	"slices"
	"strconv"
)

func Part1(input string) (string, error) {
	return "", nil
}

func joltage(bank string) (int, error) {

	l := len(bank)

	arr, err := makeArr(bank[:l-2])

	if err != nil {
		return 0, err
	}

	max := slices.Max(arr)

	i := slices.Index(arr, max)

	if i == l-1 {
		return strconv.Atoi(string(arr[l-2]) + string(arr[l]))
	}

}

func makeArr(str string) ([]int, error) {
	l := len(str)
	arr := make([]int, l)

	for i := 0; i < l; i++ {
		n, err := strconv.Atoi(string(str[i]))

		if err != nil {
			return []int{}, err
		}

		arr[i] = n
	}

	return arr, nil
}
