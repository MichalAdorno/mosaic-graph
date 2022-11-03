package main

import "testing"

func Test_GetDefaultOutputFileName(t *testing.T) {
	var tests = []struct {
		path     string
		expected string
	}{
		{
			path: "input/data.csv", expected: "out_data.png",
		},
		{
			path: "c/windows/users/user/xyz/documents/input/data.csv", expected: "out_data.png",
		},
	}
	for _, test := range tests {
		if result := getDefaultOutputFileName(test.path); result != test.expected {
			t.Errorf("getDefaultOutputFileName ---> expected %v but got %v", test.expected, result)
		}
	}
}
