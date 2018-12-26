package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPart1 = []struct {
	testName string
	in       string
	workers  int
	execTime int
	expOut1  string
	expOut2  int
}{
	{
		testName: "Normal slim",
		in:       "test_input",
		workers:  2,
		execTime: 0,
		expOut1:  "CABDFE",
		expOut2:  15,
	},
	{
		testName: "Normal",
		in:       "input",
		workers:  5,
		execTime: 60,
		expOut1:  "FMOXCDGJRAUIHKNYZTESWLPBQV",
		expOut2:  1053,
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

// CAFBD <- F too early, missing E
