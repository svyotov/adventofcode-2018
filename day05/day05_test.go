package day05

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
		expOut1:  10,
		expOut2:  4,
	},
	{
		testName: "Normal",
		in:       "input",
		expOut1:  11590,
		expOut2:  4504,
	},
}

func TestPart1(t *testing.T) {
	for _, tt := range testPart1 {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut1, Run1(tt.in), tt.testName)
			assert.Equal(t, tt.expOut2, Run2(tt.in), tt.testName)
		})
	}
}
