package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func merge(left []int, right []int) []int {

	result := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] > right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	middle := len(input) / 2

	left := MergeSort(input[:middle])
	right := MergeSort(input[middle:])

	return merge(left, right)
}

func parseIntSliceFromArgs() ([]int, error) {
	if len(os.Args) == 0 {
		return nil, errors.New("Empty data")
	}

	arr := os.Args[1:]
	data := []int{}

	for _, s := range arr {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		data = append(data, n)
	}

	return data, nil
}

func main() {

	data, err := parseIntSliceFromArgs()
	if err != nil {
		log.Fatal(err)
	}

	data = MergeSort(data)
	writeData(data)
}

func writeData(data []int) error {
	file, err := os.Create("result.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, n := range data {
		file.WriteString(fmt.Sprintf("%d\n", n))
	}

	return nil
}
