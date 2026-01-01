package runner

import "fmt"

type SolutionFn func(input string) (string, error)

type Key struct {
	Day  int
	Part int
}

type Registry map[Key]SolutionFn

func (r Registry) Add(day, part int, fn SolutionFn) {
	r[Key{Day: day, Part: part}] = fn
}

func (r Registry) Run(day, part int, input string) ( error) {
	fn, ok := r[Key{Day: day, Part: part}]
	if (!ok) {
		return fmt.Errorf("no solution registered for day %d part %d", day, part)
	} 

	sol, err := fn(input)

	if (err != nil) {
		return err
	}

	fmt.Printf("the solution for day %d part %d is %s\n", day, part, sol)

	return nil
}