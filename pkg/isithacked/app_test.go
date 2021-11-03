package isithacked

import (
	"reflect"
	"testing"
)

// Pointless test case
func TestIsItHacked(t *testing.T) {
	testTable := []struct {
		input    string
		expected []*Output
	}{
		{
			input:    "example.com",
			expected: []*Output{},
		},
		{
			input:    "www.000webhost.com",
			expected: []*Output{},
		},
	}

	for _, testCase := range testTable {
		result, _ := IsItHacked(testCase.input)
		if !(reflect.TypeOf(result) == reflect.TypeOf(testCase.expected)) {
			t.Errorf("Unexpected result! Ecpected %v - got %v", testCase.expected, result)
		}
	}
}
