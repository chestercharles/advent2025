package main

import (
	"flag"
	"log"

	"github.com/chestercharles/advent2025/internal/input"
	"github.com/chestercharles/advent2025/internal/runner"
	"github.com/chestercharles/advent2025/puzzles/day1"
)



func main() {
   var day int
   var part int
   var all bool

   flag.IntVar(&day, "day", 0, "day")
   flag.IntVar(&part, "part", 0, "part")
   flag.BoolVar(&all, "all", false, "all")
   flag.Parse()

   reg := runner.Registry{}
   registerSolutions(reg)

   if (day == 0 || part == 0 ) {
	log.Fatal("Must specify day or part")
   }

   in, err := input.Read(day)

   if (err != nil) {
	log.Fatal(err)
   }

   err2 := reg.Run(day, part, in)

   if (err2 != nil) {
	log.Fatal(err2)
   }

}

func registerSolutions(r runner.Registry)  {
	 r.Add(1, 1, day1.Part1)
}