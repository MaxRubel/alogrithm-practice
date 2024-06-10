package sortAlgos

import "errors"

func BinarySearch(arr []int, value int) (int, error) {
	first := 0
	last := len(arr) - 1
	middle := (first + last) / 2

	for first <= last {
		result := arr[middle]

		if result == value {
			return middle, nil
		}
		if value > result {
			first = middle + 1
			middle = (first + last) / 2
		} else {
			last = middle - 1
			middle = (first + last) / 2
		}
	}
	return -1, errors.New("index not found in array")
}
