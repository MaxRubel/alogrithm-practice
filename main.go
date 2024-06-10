package main

import (
	"fmt"

	sortAlgos "github.com/MaxRubel/alogrithm-practice/sort"
)

func main() {
	fmt.Println("running sorts")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	value := 15
	result, err := sortAlgos.BinarySearch(arr, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("index: ", result)
}
