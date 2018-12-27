package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPart1 = []struct {
	testName string
	players  int
	marbles  int
	expOut1  int
	expOut2  int
}{
	{
		testName: "T0",
		players:  9,
		marbles:  25,
		expOut1:  32,
		expOut2:  19708,
	},
	{
		testName: "T1",
		players:  10,
		marbles:  1618,
		expOut1:  8317,
		expOut2:  68161699,
	},
	{
		testName: "T2",
		players:  13,
		marbles:  7999,
		expOut1:  146373,
		expOut2:  1406506154,
	},
}

func Test(t *testing.T) {
	for _, tt := range testPart1 {
		t.Run(tt.testName, func(t *testing.T) {
			// assert.Equal(t, tt.expOut1, Run0(tt.players, tt.marbles), tt.testName)
			assert.Equal(t, tt.expOut1, Run1(tt.players, tt.marbles), tt.testName)
			assert.Equal(t, tt.expOut2, Run2(tt.players, tt.marbles), tt.testName)
		})
	}
}

// // CAFBD <- F too early, missing E
