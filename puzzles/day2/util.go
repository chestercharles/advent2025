package day2

import "strconv"

func makeArr(start, finish string) ([]int, error) {
	s, err := strconv.Atoi(start)
	if err != nil {
		return []int{}, err
	}

	f, err := strconv.Atoi(finish)
	if err != nil {
		return []int{}, err
	}

	size := f - s + 1
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = s + i
	}

	return arr, nil
}
