package main

import (
	"github.com/svyotov/adventofcode-2018/day03"
	"github.com/svyotov/adventofcode-2018/day04"
	"github.com/svyotov/adventofcode-2018/day05"
	"github.com/svyotov/adventofcode-2018/day06"
)

func main() {
	day03.Run1("day03/input", 1000)
	day03.Run2("day03/input", 1000)
	day04.Run1("day04/input")
	day04.Run2("day04/input")
	day05.Run1("day05/input")
	day05.Run2("day05/input")
	day06.Run1("day06/input")
	day06.Run2("day06/input", 10000)
}
