package main

import (
	"fmt"
	"strings"
)


// Given an array of strings interspersed with empty strings, search a 
// particular string in the list

func search(arr []string, start int, end int, value string) int {
	mid := (start + end) / 2
	//fmt.Println("Midpoint :", mid)
	if arr[mid] == "" {
		left := mid - 1
		right := mid + 1
		for {
			if left < start && right > end {
				return -1
			} else if left >= start && arr[left] != "" {
				mid = left
				break
			} else if right <= end && arr[right] != "" {
				mid = right
				break
			}
			right++
			left--
		}
	}
	//fmt.Println("Mid ", mid)
	diff := strings.Compare(value, arr[mid])

	if diff == 0 {
		return mid
	} else if diff > 0 {
		//fmt.Println("Searching right as start and end are: ", mid+1, end)
		return search(arr, mid+1, end, value)
	} else {
		//fmt.Println("Searching left as start and end are: ", start, mid-1)
		return search(arr, start, mid-1, value)
	}
}

func main() {
	strs := []string{"blue", "", "", "orange", "", "", "", "red", "", "", "", "yellow"}

	fmt.Println(search(strs, 0, len(strs)-1, "orange"))
	fmt.Println(search(strs, 0, len(strs)-1, "blue"))
	fmt.Println(search(strs, 0, len(strs)-1, "yellow"))
}

