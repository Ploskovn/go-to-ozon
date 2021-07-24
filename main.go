package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 77, 78, 89, 94, 99}
	subarray := []int{78, 89}

	res := isInclude(array, subarray)

	fmt.Printf("result: %t\n", res)
}

func isInclude(array []int, subarray []int) bool {
	if len(subarray) == 0 {
		return true
	}

	first := 0
	last := len(array) - 1

	for first <= last {
		mid := (first + last) / 2
		guess := array[mid]
		subFirst := subarray[0]

		if guess == subFirst {
			subLastIdx := len(subarray) - 1
			subLast := subarray[subLastIdx]

			if subLast != array[mid+subLastIdx] {
				return false
			}

			for idx, val := range subarray {
				if val != array[mid+idx] {
					return false
				}

			}

			return true

		}

		if guess > subFirst {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	return false
}
