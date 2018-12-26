package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testComputeFrequency = []struct {
	testName string
	inintial int64
	changes  []int
	expOut   int64
}{
	{testName: "No data", expOut: 0},
	{testName: "Normal starting from 0", changes: []int{1, 3, -4, 1, -10001, 2}, expOut: -9998},
	{testName: "Normal starting from -10", inintial: -10, changes: []int{1, 3, -4, 1, -10001, 2}, expOut: -10008},
}

func TestComputeFrequency(t *testing.T) {
	for _, tt := range testComputeFrequency {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut, ComputeFrequency(tt.inintial, tt.changes), tt.testName)
		})
	}
}

var testreadInputData = []struct {
	testName string
	file     string
	expLen   int
	expOut   int64
}{
	{testName: "Normal", file: "input", expLen: 993, expOut: 556},
}

func TestreadInputData(t *testing.T) {
	for _, tt := range testreadInputData {
		t.Run(tt.testName, func(t *testing.T) {
			changes, err := readInputData(tt.file)
			assert.NoError(t, err, tt.testName)
			assert.Equal(t, tt.expLen, len(changes), tt.testName)
			assert.Equal(t, tt.expOut, ComputeFrequency(0, changes), tt.testName)
		})
	}
}

var testFindFirstDuplicate = []struct {
	testName string
	inintial int64
	changes  []int
	maxLoops int
	expOut   int64
}{
	{testName: "Test 1", changes: []int{+1, -1}, expOut: 0},
	{testName: "Test 2", changes: []int{+3, +3, +4, -2, -4}, expOut: 10},
	{testName: "Test 3", changes: []int{-6, +3, +8, +5, -6}, expOut: 5},
	{testName: "Test 4", changes: []int{+7, +7, -2, -7, -4}, expOut: 14},
}

func TestFindFirstDuplicate(t *testing.T) {
	for _, tt := range testFindFirstDuplicate {
		t.Run(tt.testName, func(t *testing.T) {
			firstDuplicate, err := FindFirstDuplicate(tt.inintial, tt.changes, 10)
			assert.NoError(t, err, tt.testName)
			assert.Equal(t, tt.expOut, firstDuplicate, tt.testName)
		})
	}
}

func TestFindFirstDuplicateFile(t *testing.T) {
	testName := "from file"
	t.Run(testName, func(t *testing.T) {
		changes, err := readInputData("input")
		assert.NoError(t, err, testName)
		firstDuplicate, err := FindFirstDuplicate(0, changes, 300)
		assert.NoError(t, err, testName)
		assert.Equal(t, int64(448), firstDuplicate, testName)
	})
}
