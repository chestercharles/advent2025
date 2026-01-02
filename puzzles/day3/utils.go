package day3

import "strconv"

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
