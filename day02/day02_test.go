package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFreqCounter = []struct {
	testName string
	in       string
	expOut   map[rune]int
}{
	{testName: "T0", expOut: map[rune]int{}},
	{testName: "T1", in: "abcdef", expOut: map[rune]int{97: 1, 98: 1, 99: 1, 100: 1, 101: 1, 102: 1}},
	{testName: "T2", in: "bababc", expOut: map[rune]int{97: 2, 98: 3, 99: 1}},
}

func TestFreqCounter(t *testing.T) {
	for _, tt := range testFreqCounter {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut, freqCounter(tt.in), tt.testName)
		})
	}
}

var testCheckSum = []struct {
	testName string
	in       []string
	expOut   int
}{
	{testName: "No input", expOut: 0},
	{testName: "Normal", in: []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, expOut: 12},
}

func TestCheckSum(t *testing.T) {
	for _, tt := range testCheckSum {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut, CheckSum(tt.in), tt.testName)
		})
	}
}

func TestLocalFileHash(t *testing.T) {
	testName := "case 1"
	t.Run(testName, func(t *testing.T) {
		data, err := readInputData("input")
		if err != nil {
			panic("failed to read data " + err.Error())
		}
		assert.Equal(t, 8296, CheckSum(data), testName)
	})
}

var testFindBoxDistanceOne = []struct {
	testName string
	in       []string
	expOk    bool
	expOut   string
}{
	{testName: "No input", expOut: ""},
	{testName: "Normal", in: []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, expOk: true, expOut: "fgij"},
}

func TestFindBoxDistanceOne(t *testing.T) {
	for _, tt := range testFindBoxDistanceOne {
		t.Run(tt.testName, func(t *testing.T) {
			value, ok := FindBoxDistanceOne(tt.in)
			if tt.expOk {
				assert.True(t, ok, tt.testName)
				assert.Equal(t, tt.expOut, value, tt.testName)
			} else {
				assert.False(t, ok, tt.testName)
			}
		})
	}
}

func TestFindBoxDistanceOneFile(t *testing.T) {
	testName := "case 1"
	t.Run(testName, func(t *testing.T) {
		data, err := readInputData("input")
		if err != nil {
			panic("failed to read data " + err.Error())
		}
		match, ok := FindBoxDistanceOne(data)
		assert.True(t, ok, testName)
		assert.Equal(t, "pazvmqbftrbeosiecxlghkwud", match, testName)
	})
}
