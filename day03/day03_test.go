package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCountOverlapingSquares = []struct {
	testName string
	in       []Claim
	size     int
	expOut   int
}{
	{
		testName: "T0",
		in: []Claim{{ID: 1, X: 1, Y: 3, W: 4, H: 4},
			{ID: 2, X: 3, Y: 1, W: 4, H: 4},
			{ID: 3, X: 5, Y: 5, W: 2, H: 2}},
		size:   8,
		expOut: 4,
	},
}

func TestCountOverlapingSquares(t *testing.T) {
	for _, tt := range testCountOverlapingSquares {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut, CountOverlapingSquares(tt.size, tt.in), tt.testName)
		})
	}
}

func TestCountOverlapingSquaresFile(t *testing.T) {
	testName := "case 1"
	t.Run(testName, func(t *testing.T) {
		data, err := ReadInputData("input")
		if err != nil {
			panic("failed to read data " + err.Error())
		}
		assert.Equal(t, 113576, CountOverlapingSquares(1000, data), testName)
	})
}

var testGetNonOverlapingData = []struct {
	testName string
	in       []Claim
	expOut   Claim
}{
	{
		testName: "T0",
		in: []Claim{{ID: 1, X: 1, Y: 3, W: 4, H: 4},
			{ID: 2, X: 3, Y: 1, W: 4, H: 4},
			{ID: 3, X: 5, Y: 5, W: 2, H: 2}},
		expOut: Claim{ID: 3, X: 5, Y: 5, W: 2, H: 2},
	},
}

func TestGetNonOverlapingData(t *testing.T) {
	for _, tt := range testGetNonOverlapingData {
		t.Run(tt.testName, func(t *testing.T) {
			claim, ok := GetNonOverlapingData(1000, tt.in)
			assert.True(t, ok, tt.testName)
			assert.Equal(t, tt.expOut, claim, tt.testName)
		})
	}
}

func TestGetNonOverlapingDataFile(t *testing.T) {
	testName := "case 1"
	t.Run(testName, func(t *testing.T) {
		data, err := ReadInputData("input")
		if err != nil {
			panic("failed to read data " + err.Error())
		}

		claim, ok := GetNonOverlapingData(1000, data)
		assert.True(t, ok, testName)
		assert.Equal(t, Claim{ID: 825, X: 689, Y: 535, W: 23, H: 27}, claim, testName)
	})
}
