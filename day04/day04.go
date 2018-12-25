package day04

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	beginsShift = iota // 0
	fallsAsleep        // 1
	wakesUp            // 2
	endShift           //3
)

// event is a Guard action representation
type event struct {
	date time.Time
	text string
	// etype
	// 0 : begins shift
	// 1 : falls asleep
	// 2 : wakes up
	// 3 : ends shift
	etype int
}

// guard is used to store the  informations such as
// [1518-11-05 00:03] Guard #99 begins shift
// [1518-11-05 00:45] falls asleep
// [1518-11-05 00:55] wakes up
type guard struct {
	id     int
	events events
}

type events []event

func (e events) Len() int {
	return len(e)
}

func (e events) Less(i, j int) bool {
	return e[i].date.Before(e[j].date)
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// readInputData reads the input data for this test
// asumes no header
func readInputData(file string) (data map[int]guard, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return map[int]guard{}, err
	}

	rows := strings.Split(strings.Trim(string(b), "\n"), "\n")
	data = make(map[int]guard, len(rows))
	guardID := -1
	date, _ := time.Parse("2006-01-02 15:04", "2006-01-02 15:04")
	re := regexp.MustCompile("[0-9]+")

	allEvents := make(events, len(rows))

	for indx, row := range rows {
		date, err := time.Parse("2006-01-02 15:04", row[1:17])
		if err != nil {
			panic(err)
		}
		allEvents[indx] = event{date: date, text: row, etype: -1}
	}

	// sort the raw data based on time stamps
	sort.Sort(allEvents)

	for _, evnt := range allEvents {
		currentGuard := data[guardID]
		currentGuard.id = guardID

		switch keyword := evnt.text[19:24]; keyword {
		case "Guard":
			// close the old guard, if the first time we will just add a fake one -1
			// this is more efficient than having an if, and can be removed later
			evnt.etype = endShift
			data[guardID] = guard{id: guardID, events: append(data[guardID].events, evnt)}
			guardID, err = strconv.Atoi(re.FindAllString(evnt.text[26:], -1)[0])
			if err != nil {
				panic(err)
			}
			// start the new guard
			evnt.etype = beginsShift
			data[guardID] = guard{id: guardID, events: append(data[guardID].events, evnt)}
		case "falls":
			evnt.etype = fallsAsleep
			data[guardID] = guard{id: guardID, events: append(data[guardID].events, evnt)}
		case "wakes":
			evnt.etype = wakesUp
			data[guardID] = guard{id: guardID, events: append(data[guardID].events, evnt)}
		default:
			panic("unsupported keyword " + keyword)
		}
	}
	// remove -1 guard
	delete(data, -1)
	// close the last guard
	data[guardID] = guard{id: guardID, events: append(data[guardID].events, event{etype: endShift, date: date})}

	for gid := range data {
		sort.Sort(data[gid].events)
	}

	return data, err
}

func updateSlepTime(start, end time.Time, sleep []int) {
	if start.After(end) {
		panic("start after end")
	}
	for d := start; d.Before(end); d = d.Add(time.Minute) {
		// do stuff with d
		if d.Hour() == 0 {
			sleep[d.Minute()]++
		}
	}
}

func sum(in []int) int {
	t := 0
	for _, v := range in {
		t += int(v)
	}
	return t
}

func max(in []int) (int, int) {
	m := -1
	indx := -1
	for i, v := range in {
		if v > m {
			m = v
			indx = i
		}
	}
	return m, indx
}

func mostAsleepMinute(in []int) int {
	m := -1
	i := -1
	for k, v := range in {
		if v > m {
			m = v
			i = k
		}
	}
	return i
}

// sleepEstimatorTask1 estimates the number of hours
// slept times the guard ID for the most sleepy guard
// as described in part 1 in README.md
func sleepEstimatorTask1(data map[int]guard) int {
	maxTotalSleep := 0
	mostAsleepMinuteVal := -1
	maxTotalSleepGuardID := -1

	for _, guard := range data {
		e := guard.events
		sleep := make([]int, 60)
		for i := 0; i < len(e); i++ {
			if e[i].etype == fallsAsleep {
				if e[i+1].etype != endShift && e[i+1].etype != wakesUp {
					panic("unexpected etype" + string(e[i+1].etype))
				}
				updateSlepTime(e[i].date, e[i+1].date, sleep)
			}
		}
		tmpMaxTotalSleep := sum(sleep)
		if tmpMaxTotalSleep > maxTotalSleep {
			maxTotalSleep = tmpMaxTotalSleep
			mostAsleepMinuteVal = mostAsleepMinute(sleep)
			maxTotalSleepGuardID = guard.id
		}
	}
	return mostAsleepMinuteVal * maxTotalSleepGuardID
}

// sleepEstimatorTask1 estimates the number of hours
// slept times the guard ID for the most sleepy guard
// as described in part 1 in README.md
func sleepEstimatorTask2(data map[int]guard) int {
	maxTotalSleep := -1
	maxTotalSleepIndx := -1
	maxTotalSleepGuardID := -1

	for _, guard := range data {
		e := guard.events
		sleep := make([]int, 60)
		for i := 0; i < len(e); i++ {
			if e[i].etype == fallsAsleep {
				if e[i+1].etype != endShift && e[i+1].etype != wakesUp {
					panic("unexpected etype" + string(e[i+1].etype))
				}
				updateSlepTime(e[i].date, e[i+1].date, sleep)
			}
		}
		tmpMaxTotalSleep, tmpIndx := max(sleep)
		if tmpMaxTotalSleep > maxTotalSleep {
			maxTotalSleep = tmpMaxTotalSleep
			maxTotalSleepIndx = tmpIndx
			maxTotalSleepGuardID = guard.id
		}
	}
	return maxTotalSleepIndx * maxTotalSleepGuardID
}

// Run1 runs task one for this days ta/home/svyotov/go/src/github.com/svyotov/adventofcode-2018/day04/day04.go:234:3: syntax error: unexpected input at end of statement
func Run1(file string) {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 04 t1: '%v'\n", sleepEstimatorTask1(data))
}

// Run2 runs task tow for this days task
func Run2(file string) {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 04 t2: '%v'\n", sleepEstimatorTask2(data))
}
