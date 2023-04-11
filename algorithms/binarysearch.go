package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	ok := bs_list(arr, 2)
	if !ok {
		fmt.Println("Haven't found it!")
	}
}

func bs_list[T any](arr []int, needle ) bool {
	lo := 0
	hi := len(arr)
	for lo < hi {
		m := math.Floor(float64(lo + (hi-lo)/2))
		v := arr[m]
		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
