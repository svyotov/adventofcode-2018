package day05

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

// readInputData reads the input data for this test
// asumes no header
func readInputData(file string) (data string, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(b), "\n"), err
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

func remove(slice string, s int, e int) string {
	return slice[:s] + slice[e:]
}

// simplifiedLength simplifies the array and returns it's length
// as described in part 1 in README.md
func simplifiedLength(data string) int {
	ldata := strings.ToLower(data)
	for index := 0; index < len(ldata)-1; {
		if ldata[index] == ldata[index+1] && data[index] != data[index+1] {
			ldata = remove(ldata, index, index+2)
			data = remove(data, index, index+2)
			if index > 0 {
				index--
			}
		} else {
			index++
		}
	}
	return len(data)
}

func toCharStr(i int) string {
	return fmt.Sprintf("[%c|%c]", 'A'-1+i, 'a'-1+i)
}

func shortestPolymerLength(data string) int {
	minScore := simplifiedLength(data)
	for i := 1; i < 27; i++ {
		var re = regexp.MustCompile(toCharStr(i))
		sdata := re.ReplaceAllString(data, "")
		tmpScore := simplifiedLength(sdata)
		if tmpScore < minScore {
			minScore = tmpScore
		}
	}
	return minScore
}

// Run1 task A fr current day
func Run1(file string) int {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	score := simplifiedLength(data)
	fmt.Printf("Day 05 t1: '%v'\n", score)
	return score
}

// Run2 task B fr current day
func Run2(file string) int {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	score := shortestPolymerLength(data)
	fmt.Printf("Day 05 t1: '%v'\n", score)
	return score
}
