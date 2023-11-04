package contains

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	testTable := []struct {
		wayToFile   string
		str         string
		rightResult bool
	}{
		{
			wayToFile:   "test.txt",
			str:         "qwerty",
			rightResult: true,
		},
		{
			wayToFile:   "test.txt",
			str:         "aasd",
			rightResult: false,
		},
	}
	for _, testCase := range testTable {
		result, err := Contains(testCase.wayToFile, testCase.str)
		if err != nil {
			fmt.Println(err)
		} else {
			t.Logf("Test way to file: %s, test string for search: %s", testCase.wayToFile, testCase.str)
			if result != testCase.rightResult {
				t.Error("Incorrect result. Expect: ", testCase.rightResult, "got: ", testCase.rightResult, result)
			}
		}
	}
}
