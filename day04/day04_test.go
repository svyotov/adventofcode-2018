package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPart1 = []struct {
	testName string
	in       string
	expOut1  int
	expOut2  int
}{
	{
		testName: "Normal",
		in:       "test_input",
		expOut1:  240,
		expOut2:  4455,
	},
	{
		testName: "Normal",
		in:       "input",
		expOut1:  30630,
		expOut2:  136571,
	},
}

func TestPart1(t *testing.T) {
	for _, tt := range testPart1 {
		t.Run(tt.testName, func(t *testing.T) {
			data, err := readInputData(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.expOut1, sleepEstimatorTask1(data), tt.testName)
			assert.Equal(t, tt.expOut2, sleepEstimatorTask2(data), tt.testName)
		})
	}
}
