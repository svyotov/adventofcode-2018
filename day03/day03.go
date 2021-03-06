package day03

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Claim such as '#123 @ 3,2: 5x4' means that claim
// ID 123 specifies a rectangle
// 3 inches from the left edge, 2 inches from the top edge,
// 5 inches wide, and 4 inches tall.
type Claim struct {
	ID         int
	X, Y, W, H int
}

// ToString converts a claim to it's string representation
func (c Claim) ToString() string {
	return fmt.Sprintf("#%v @ %v,%v: %vx%v", c.ID, c.X, c.Y, c.W, c.H)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// readInputData reads the input data for this test
// adssumes no header
func readInputData(file string) (data []Claim, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return []Claim{}, err
	}

	rows := strings.Split(strings.Trim(string(b), "\n"), "\n")
	data = make([]Claim, len(rows))

	for index := range rows {
		array := strings.FieldsFunc(rows[index], split)
		id, err := strconv.Atoi(array[0])
		handleErr(err)
		x, err := strconv.Atoi(array[1])
		handleErr(err)
		y, err := strconv.Atoi(array[2])
		handleErr(err)
		w, err := strconv.Atoi(array[3])
		handleErr(err)
		h, err := strconv.Atoi(array[4])
		handleErr(err)

		data[index] = Claim{ID: id, X: x, Y: y, W: w, H: h}
	}

	return data, err
}

func split(r rune) bool {
	return r == ' ' || r == ',' || r == '#' || r == ':' || r == 'x' || r == '@'
}

// CountOverlappingSquares counts the # of squares overlapping
// as described in part 1 in README.md
// TODO this will panic if claims out of range (size)
func CountOverlappingSquares(size int, data []Claim) int {
	field := getOverlapField(size, data)

	// get the number of the overlapping fields
	sum := 0
	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field[x]); y++ {
			if field[x][y] > 1 {
				sum++
			}
		}
	}

	return sum
}

func getOverlapField(size int, data []Claim) [][]int {
	// Create two-dimensional array.
	field := make([][]int, size)
	for i := range field {
		field[i] = make([]int, size)
	}

	// find the overlapping fields
	for _, claim := range data {
		for w := 0; w < claim.W; w++ {
			for h := 0; h < claim.H; h++ {
				field[claim.X+w][claim.Y+h]++
			}

		}
	}
	return field
}

// GetNonOverlappingData counts the # of squares overlapping
// as described in part 1 in README.md
// TODO this will panic if claims out of range (size)
func GetNonOverlappingData(size int, data []Claim) (Claim, bool) {
	field := getOverlapField(size, data)

	// fiend the overlapping fields
	for _, claim := range data {
		ok := true
		for w := 0; w < claim.W; w++ {
			for h := 0; h < claim.H; h++ {
				if field[claim.X+w][claim.Y+h] != 1 {
					ok = false
				}
			}
		}
		if ok {
			return claim, true
		}
	}
	return Claim{}, false
}

// Run1 runs task one for this day
func Run1(file string, threshold int) {
	data, err := readInputData(file)
	handleErr(err)

	fmt.Printf("Day 03 t1: '%v'\n", CountOverlappingSquares(threshold, data))
}

// Run2 runs task two for this day
func Run2(file string, threshold int) {
	data, err := readInputData(file)
	handleErr(err)
	claim, ok := GetNonOverlappingData(threshold, data)
	if !ok {
		panic("failed to run day 3 task 2")
	}
	fmt.Printf("Day 03 t2: '%v'\n", claim.ID)
}
