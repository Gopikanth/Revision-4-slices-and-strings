//slices are flexible and light weight data structure which stores similar type of datas
package main

import "fmt"

func main() {

	//creating slices
	new_slice_1 := []string{"i", "love", "india"}
	fmt.Println(new_slice_1)
	//creating slices from array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	new_slice := arr[1:4]
	fmt.Println(new_slice)
	fmt.Println(len(new_slice)) //length of slices
	fmt.Println(cap(new_slice)) //capacity of slices
	//creating slices from slices
	new_slice_3 := []int{1,2,3,4,5,6,7,8}
	first := new_slice_3[0:3]
	second := new_slice_3[3:5]
	third  := new_slice_3[4:] 
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	//creating slice using make functions
	new_slice_4 := make([]int,4,5)
	fmt.Println(len(new_slice_4))
	fmt.Println(cap(new_slice_4))
	

}
