package main

import "testing"

type testdata struct {
	array    []int
	subarray []int
	res      bool
}

var tests = []testdata{
	{[]int{1, 2, 3, 5, 7, 9, 11}, []int{}, true},
	{[]int{1, 2, 3, 5, 7, 9, 11}, []int{3, 5, 7}, true},
	{[]int{1, 2, 3, 5, 7, 9, 11}, []int{4, 5, 7}, false},
}

func TestIsInclude(t *testing.T) {
	for _, data := range tests {
		result := isInclude(data.array, data.subarray)
		if result != data.res {
			t.Error(
				"For", data.array,
				"and", data.subarray,
				"expected", data.res,
				"got", result,
			)
		}
	}
}
