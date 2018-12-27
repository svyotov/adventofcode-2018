package day06

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const infinite = "infinite"

// point represents (x,y) coordinates
type point struct {
	id   string
	x, y int
}

// fpoint represents a point on the field
type fpoint struct {
	owner string
	score int
}

// readInputData reads the input data for this test
// asumes no header
func readInputData(file string) (data []point, err error) {
	f, err := os.Open(file)
	if err != nil {
		return data, err
	}

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		return data, err
	}

	data = make([]point, len(lines))

	for i, line := range lines {
		data[i].y, err = strconv.Atoi(line[0])
		if err != nil {
			return data, err
		}
		data[i].x, err = strconv.Atoi(strings.TrimSpace(line[1]))
		if err != nil {
			return data, err
		}
		data[i].id = fmt.Sprintf("#%v", i)
	}
	return data, err
}

func generateField(a point, fx, fy int) [][]fpoint {
	field := make([][]fpoint, fx)
	for ix := 0; ix < fx; ix++ {
		field[ix] = make([]fpoint, fy)
	}
	for ix := 0; ix < fx; ix++ {
		for iy := 0; iy < fy; iy++ {
			field[ix][iy].owner = a.id
		}
	}

	// first square
	for ix := 0; ix < a.x; ix++ {
		xScore := a.x - ix
		for iy := 0; iy < a.y; iy++ {
			field[ix][iy].score = xScore + a.y - iy
		}
	}

	// second square
	for ix := a.x; ix < fx; ix++ {
		xScore := ix - a.x
		for iy := 0; iy < a.y; iy++ {
			field[ix][iy].score = xScore + a.y - iy
		}
	}

	// third square
	for ix := 0; ix < a.x; ix++ {
		xScore := a.x - ix
		for iy := a.y; iy < fy; iy++ {
			field[ix][iy].score = xScore + iy - a.y
		}
	}

	// last square
	for ix := a.x; ix < fx; ix++ {
		xScore := ix - a.x
		for iy := a.y; iy < fy; iy++ {
			field[ix][iy].score = xScore + iy - a.y
		}
	}

	return field
}

// merge will merge two fileds
// resulting in a representation
// like bellow, where all the lowecase
// letters are the manhatan distance of
// corresponding the upercase, and '.' is infinity
//
// aaaaa.cccc
// aAaaa.cccc
// aaaddecccc
// aadddeccCc
// ..dDdeeccc
// bb.deEeecc
// bBb.eeee..
// bbb.eeefff
// bbb.eeffff
// bbb.ffffFf
func merge(a, b [][]fpoint) [][]fpoint {
	result := copy(a)
	for ix := range a {
		for iy := range a[ix] {
			if a[ix][iy].score == b[ix][iy].score {
				result[ix][iy].owner = infinite
			}
			if a[ix][iy].score > b[ix][iy].score {
				result[ix][iy] = b[ix][iy]
			}
		}
	}
	return result
}

func copy(slice [][]fpoint) [][]fpoint {
	cp := make([][]fpoint, len(slice))
	for i := range slice {
		cp[i] = append([]fpoint{}, slice[i]...)
	}
	return cp
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func getXY(data []point) (int, int) {
	x, y := 0, 0
	for _, p := range data {
		x = max(x, p.x)
		y = max(y, p.y)
	}
	N := max(x, y) + 1
	return N, N
}

func getMaxScore(data [][]fpoint) int {
	score := 0
	for _, row := range data {
		for _, p := range row {
			score = max(score, p.score)
		}
	}
	return score
}

// simplifiedLength simplifies the array and returns it's length
// as described in part 1 in README.md
func getGlobalField(data []point) [][]fpoint {
	x, y := getXY(data)
	field := generateField(data[0], x, y)
	for _, p := range data[1:] {
		tf := generateField(p, x, y)
		field = merge(tf, field)
	}
	return field
}

func print(field [][]fpoint) {
	for x := range field {
		for y := range field[x] {
			fmt.Print(field[x][y].score)
		}
		fmt.Println()
	}
	fmt.Println("----------")
	for x := range field {
		for y := range field[x] {
			fmt.Printf(field[x][y].owner[1:2])
		}
		fmt.Println()
	}
	fmt.Println("===========")
}

type score struct {
	border bool
	score  int
}

func getMaxField(data []point) int {
	field := getGlobalField(data)
	points := map[string]score{}
	for _, row := range field {
		for _, dt := range row {
			pnt := points[dt.owner]
			pnt.score++
			points[dt.owner] = pnt
		}
	}

	nx := len(field)
	ny := len(field[0])
	for indx := 0; indx < nx; indx++ {
		id := field[indx][0].owner
		pnt := points[id]
		pnt.border = true
		points[id] = pnt
		//
		id = field[indx][ny-1].owner
		pnt = points[id]
		pnt.border = true
		points[id] = pnt
	}

	for indx := 0; indx < ny; indx++ {
		id := field[0][indx].owner
		pnt := points[id]
		pnt.border = true
		points[id] = pnt
		//
		id = field[nx-1][indx].owner
		pnt = points[id]
		pnt.border = true
		points[id] = pnt
	}
	maxScore := 0
	for _, p := range points {
		if !p.border && p.score > maxScore {
			maxScore = p.score
		}
	}

	return maxScore
}

// Run1 task A for current day
func Run1(file string) int {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	score := getMaxField(data)
	fmt.Printf("Day 06 t1: '%v'\n", score)
	return score
}

// simplifiedLength simplifies the array and returns it's length
// as described in part 2 in README.md
func getGlobalField2(data []point) [][]fpoint {
	x, y := getXY(data)
	field := generateField(data[0], x, y)
	for _, p := range data[1:] {
		tf := generateField(p, x, y)
		field = merge2(tf, field)
	}
	return field
}

func merge2(a, b [][]fpoint) [][]fpoint {
	result := copy(a)
	for ix := range a {
		for iy := range a[ix] {
			result[ix][iy].score = a[ix][iy].score + b[ix][iy].score
		}
	}
	return result
}

func getFieldsUnderTreshold(data []point, threshold int) int {
	field := getGlobalField2(data)
	count := 0
	for _, row := range field {
		for _, dt := range row {
			if dt.score < threshold {
				count++
			}
		}
	}
	return count
}

// Run2 task B for current day
func Run2(file string, threshold int) int {
	data, err := readInputData(file)
	if err != nil {
		panic(err)
	}
	score := getFieldsUnderTreshold(data, threshold)
	fmt.Printf("Day 06 t2: '%v'\n", score)
	return score
}
