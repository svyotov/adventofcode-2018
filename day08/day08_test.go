package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPart1 = []struct {
	testName string
	in       string
	workers  int
	execTime int
	expOut1  int
	expOut2  int
}{
	{
		testName: "Normal slim",
		in:       "test_input",
		workers:  2,
		execTime: 0,
		expOut1:  138,
		expOut2:  66,
	},
	{
		testName: "Normal",
		in:       "input",
		workers:  5,
		execTime: 60,
		expOut1:  40036,
		expOut2:  21677,
	},
}

func Test(t *testing.T) {
	for _, tt := range testPart1 {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expOut1, Run1(tt.in), tt.testName)
			assert.Equal(t, tt.expOut2, Run2(tt.in, tt.workers, tt.execTime), tt.testName)
		})
	}
}

// // CAFBD <- F too early, missing E
