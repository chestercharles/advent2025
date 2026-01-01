package input

import (
	"fmt"
	"os"
)

func Read(day int) (string, error) {
	path := fmt.Sprintf("inputs/day%d.txt", day)
	b, err := os.ReadFile(path)
	if (err != nil) {
		return "", err
	}

	return string(b), nil
}