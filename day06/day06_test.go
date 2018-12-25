package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPart1 = []struct {
	testName  string
	in        string
	threshold int
	expOut1   int
	expOut2   int
}{
	{
		testName:  "Normal slim",
		in:        "test_input",
		threshold: 32,
		expOut1:   17,
		expOut2:   16,
	},
	{
		testName:  "Normal",
		in:        "input",
		threshold: 10000,
		expOut1:   4976,
		expOut2:   46462,
	},
}

func TestPart1(t *testing.T) {
	for _, tt := range testPart1 {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut1, Run1(tt.in), tt.testName)
			assert.Equal(t, tt.expOut2, Run2(tt.in, tt.threshold), tt.testName)
		})
	}
}
