package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{2, 7, 9, 11, 15, 6}
	target := 9
	sum := TwoSumNaive(slice, target)
	fmt.Println(sum)
}

//bruteforce implementation
func TwoSumNaive(slice []int, target int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := i; j < len(slice); j++ {
			sum := slice[i] + slice[j]
			if sum == target {
				return []int{i, j}
			}
			continue
		}
	}
	return []int{}
}

// using HashMap
func TwoSumHashmap(slice []int, target int) []int {
	hashMap := map[int]int{}
	for idx := 0; idx < len(slice); idx++ {
		pmatch := target - slice[idx]
		if v, exists := hashMap[pmatch]; exists {
			return []int{idx, v}
		}
		hashMap[slice[idx]] = idx
	}
	return nil
}

// using two pointers
func TwoSumSortdArray(slice []int, target int) []int {
	if len(slice) < 2 {
		fmt.Println("Can't process")
		return nil
	}

	sort.Ints(slice)
	start := 0
	end := len(slice) - 1
	fmt.Println("After sorting:", slice)

	if start != end {
		sum := slice[start] + slice[end]
		if sum > target {
			end = end - 1
		} else if sum < target {
			start = start + 1
		} else {
			return []int{start, end}
		}
	}
	return nil
}
