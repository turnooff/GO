package function

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	testTable := []struct {
		n           int64
		flag        bool
		rightResult int64
	}{
		{
			n:           6,
			flag:        true,
			rightResult: 720,
		},
	}
	for _, testCase := range testTable {
		result := Calculate(testCase.n, testCase.flag)
		t.Logf("Test n: %d, test flag: %t, Result: %d", testCase.n, testCase.flag, testCase.rightResult)
		if result != testCase.rightResult {
			t.Error("Incorrect result. Expect: ", testCase.rightResult, "got: ", testCase.rightResult, result)
		}
	}
}
