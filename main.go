package main

import (
	"github.com/svyotov/adventofcode-2018/day01"
	"github.com/svyotov/adventofcode-2018/day02"
	"github.com/svyotov/adventofcode-2018/day03"
	"github.com/svyotov/adventofcode-2018/day04"
	"github.com/svyotov/adventofcode-2018/day05"
	"github.com/svyotov/adventofcode-2018/day06"
	"github.com/svyotov/adventofcode-2018/day07"
	"github.com/svyotov/adventofcode-2018/day08"
	"github.com/svyotov/adventofcode-2018/day09"
)

func main() {
	day01.Run1("day01/input")
	day01.Run2("day01/input")
	day02.Run1("day02/input")
	day02.Run2("day02/input")
	day03.Run1("day03/input", 1000)
	day03.Run2("day03/input", 1000)
	day04.Run1("day04/input")
	day04.Run2("day04/input")
	day05.Run1("day05/input")
	day05.Run2("day05/input")
	day06.Run1("day06/input")
	day06.Run2("day06/input", 10000)
	day07.Run1("day07/input")
	day07.Run2("day07/input", 5, 60)
	day08.Run1("day08/input")
	day08.Run2("day08/input", 5, 60)
	day09.Run1(418, 70769)
	day09.Run2(418, 70769)
}
