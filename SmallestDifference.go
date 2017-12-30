package main

import (
	"fmt"
	"math"
	"sort"
)

// Given two arrays, compute pair the pair of values
// (one value in each array) with the smallest (non-negative)
// difference. Return the difference

func BinarySearch(a []float64, l int, r int,
	key float64, diff float64) float64 {
	if l >= r {
		new_diff := math.Abs(key - a[l])
		if diff >= new_diff {
			return new_diff
		} else {
			return diff
		}
	}
	mid := (l + r) / 2
	d := math.Abs(key - a[mid])
	if key < a[mid] {
		return (BinarySearch(a, l, mid-1, key, d))
	} else if key > a[mid] {
		return (BinarySearch(a, mid+1, r, key, d))
	} else if key == a[mid] {
		return 0
	}
	return -1
}

const MaxInt uint64 = 1<<64 - 1

func main() {
	a := []float64{1, 3, 15, 11, 2}
	b := []float64{23, 127, 235, 19, 8}
	sort.Float64s(b)
	fmt.Println(b)
	min := float64(MaxInt)
	for _, j := range a {
		temp := BinarySearch(b, 0, len(b)-1, j, 0)
		if temp < min {
			min = temp
		}
	}
	fmt.Println("Minimum pair :", min)
}
