package contains

import (
	"os"
)

func searchStr(content string, startStr string) bool {
	for i := 0; i < len(content); i++ {
		var helpLen int = 0
		var helpI int = i
		for j := 0; j < len(startStr); j++ {
			if startStr[j] != content[i] {
				i = helpI
				break
			} else {
				helpLen++
				i++
				if helpLen == len(startStr) && j == (len(startStr)-1) {
					return true
				}
			}
		}
	}
	return false
}

func Contains(wayToFile string, str string) (bool, error) {
	content, err := os.ReadFile(wayToFile)
	if err != nil {
		return false, err
	}
	return searchStr(string(content), str), err
}
