package isithacked

import (
	"github.com/fabelx/isithacked/pkg/config"
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

func TestSelectRegex(t *testing.T) {
	testTable := []struct {
		input    bool
		expected string
	}{
		{
			input:    true,
			expected: config.IpRegex,
		},
		{
			input:    false,
			expected: config.DomainRegex,
		},
	}

	for _, testCase := range testTable {
		result := selectRegex(testCase.input)

		t.Logf("Calling selectRegex(%v), result %v", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Unexpected result! Ecpected %v - got %v", testCase.expected, result)
		}
	}
}
