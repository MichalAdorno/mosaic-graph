package main

import (
	"testing"
)

func Test_SplitIntoTwo(t *testing.T) {
	var tests = []struct {
		data  []DataPair
		left  []DataPair
		right []DataPair
		err   error
	}{
		{
			data:  []DataPair{{"A", 1}, {"B", 1}, {"C", 1}},
			left:  []DataPair{{"A", 1}},
			right: []DataPair{{"B", 1}, {"C", 1}},
			err:   nil,
		},
		{
			data:  []DataPair{{"A", 1}, {"B", 1}, {"C", 1}, {"D", 1}},
			left:  []DataPair{{"A", 1}, {"B", 1}},
			right: []DataPair{{"C", 1}, {"D", 1}},
			err:   nil,
		},
		{
			data:  []DataPair{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}},
			left:  []DataPair{{"A", 10}, {"B", 20}},
			right: []DataPair{{"C", 30}, {"D", 40}},
			err:   nil,
		},
		{
			data:  []DataPair{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}, {"E", 100}},
			left:  []DataPair{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}},
			right: []DataPair{{"E", 100}},
			err:   nil,
		},
		{
			data:  []DataPair{{"A", 10}},
			left:  []DataPair{{"A", 10}},
			right: nil,
			err:   nil,
		},
		{
			data:  []DataPair{{"AA", 1000}, {"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}, {"E", 100}},
			left:  []DataPair{{"AA", 1000}},
			right: []DataPair{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}, {"E", 100}},
			err:   nil,
		},
	}
	for _, test := range tests {
		if left, right, err := splitIntoTwo(test.data); !isEq(test.left, left) || !isEq(test.right, right) || test.err != err {
			if !isEq(test.left, left) {
				t.Errorf("splitIntoTwo ---> LEFT is: expected %v but got %v", test.left, left)
			}
			if !isEq(test.right, right) {
				t.Errorf("splitIntoTwo ---> RIGHT is: expected %v but got %v", test.right, right)
			}
			t.Errorf("-------")
		}
	}
}

func isEq(a, b []DataPair) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {

		if a[i].category != b[i].category || a[i].value != b[i].value {
			return false
		}
	}
	return true
}
