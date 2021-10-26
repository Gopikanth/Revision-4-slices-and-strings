package main

import (
	"fmt"
	"sort"
)

//zero value slice
func main() {
	new_slice_6 := []int{}
	fmt.Println(len(new_slice_6))
	fmt.Println(cap(new_slice_6))
	//slice is of reference type so it affects the array
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	new_slice_7 := arr[0:2]
	new_slice_7[0] = 4
	fmt.Println(arr)
	//comparing the slices if it is nil
	//we cant compare two slices
	fmt.Println(new_slice_6 == nil)
	//fmt.Println(new_slice_7==new_slice_6)- it throws error

	//multi-dimensional slices
	new_slice_8 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(new_slice_8)

	//sorting of slices
	new_slice_9 := []int{78, 45, 43, 29, 83, 56, 92, 95}
	fmt.Printf("The values before sorting : %v\n", new_slice_9)
	sort.Ints(new_slice_9)
	fmt.Printf("The values after sorting : %v\n", new_slice_9)
	//sorting are done in ascending order

	//checking wheather the slices are sorted

	check := sort.IntsAreSorted(new_slice_9) //returns true or false
	fmt.Println(check)
}
