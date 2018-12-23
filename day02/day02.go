package day02

import (
	"encoding/csv"
	"os"
)

// ReadInputData reads the input data for this test
// adssumes no header
func ReadInputData(file string) (data []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return data, err
	}

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		return data, err
	}

	data = make([]string, len(lines))

	for i, line := range lines {
		data[i] = line[0]
	}
	return data, err
}

// CheckSum implements a simple checksum
// as described in README.md
func CheckSum(boxIDs []string) int {
	counterOfTwo := 0
	counterOfThree := 0
	for _, boxID := range boxIDs {
		freq := freqCounter(boxID)
		if contains(freq, 2) {
			counterOfTwo++
		}
		if contains(freq, 3) {
			counterOfThree++
		}
	}
	return counterOfTwo * counterOfThree
}

// freqCounter computes the char frequencies
// for a given string
func freqCounter(boxIDs string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range boxIDs {
		freq[char]++
	}
	return freq
}

func contains(slice map[rune]int, value int) bool {
	for _, svalue := range slice {
		if svalue == value {
			return true
		}
	}
	return false
}

// FindBoxDistanceOne solves day two part 1
func FindBoxDistanceOne(boxIDs []string) (string, bool) {
	N := len(boxIDs)
	for index1 := 0; index1 < N-1; index1++ {
		for index2 := index1 + 1; index2 < N; index2++ {
			if match, ok := matchingStrings(boxIDs[index1], boxIDs[index2]); ok {
				return match, ok
			}
		}
	}
	return "", false
}
func matchingStrings(a, b string) (string, bool) {
	if simpledistance(a, b) == 1 {
		return overlapOfDistanceOne(a, b), true
	}
	return "", false
}
func simpledistance(x, y string) int {
	Nx := len(x)
	Ny := len(y)
	N := 0
	distance := 0
	if Nx >= Ny {
		N = Ny
		distance = Nx - Ny
	} else {
		N = Nx
		distance = Ny - Nx
	}

	for index := 0; index < N; index++ {
		if x[index] != y[index] {
			distance++
		}
	}
	return distance
}

func overlapOfDistanceOne(x, y string) string {
	N := min(len(x), len(y))
	for index := 0; index < N; index++ {
		if x[index] != y[index] {
			return x[0:index] + x[index+1:N]
		}
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
