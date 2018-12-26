package day01

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// readInputData reads the input data for this test
// adssumes no header
func readInputData(file string) (data []int, err error) {
	f, err := os.Open(file)
	if err != nil {
		return data, err
	}

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		return data, err
	}

	data = make([]int, len(lines))

	for i, line := range lines {
		data[i], err = strconv.Atoi(line[0])
		if err != nil {
			return data, err
		}
	}
	return data, err
}

// ComputeFrequency implements the frequency compute
// as described in README.md
func ComputeFrequency(freq int64, changes []int) int64 {
	for i := range changes {
		freq += int64(changes[i])
	}
	return freq
}

// FindFirstDuplicate implements part two of the challenge
func FindFirstDuplicate(freq int64, changes []int, maxIterations int) (int64, error) {
	previousFreq := make(map[int64]bool)
	previousFreq[freq] = true
	for iter := 0; iter < maxIterations; iter++ {
		for ci := range changes {
			freq += int64(changes[ci])
			if previousFreq[freq] {
				return freq, nil
			}
			previousFreq[freq] = true
		}
	}
	return 0, fmt.Errorf("failed to find a duplicate in %v iterations", maxIterations)
}

// Run1 runs task one for this day
func Run1(file string) {
	changes, err := readInputData(file)
	if err != nil {
		panic("failed to read data " + err.Error())
	}
	fmt.Printf("Day 01 t1: '%v'\n", ComputeFrequency(0, changes)) // 556
}

// Run2 runs task two for this day
func Run2(file string) {
	changes, err := readInputData(file)
	if err != nil {
		panic("failed to read data " + err.Error())
	}
	firstDuplicate, err := FindFirstDuplicate(0, changes, 300)
	if err != nil {
		panic("failed to read data " + err.Error())
	}

	fmt.Printf("Day 01 t1: '%v'\n", firstDuplicate) // 448
}
