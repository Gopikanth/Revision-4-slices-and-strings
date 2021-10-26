package main

import "fmt"

func main() {
	//iteration over for loop
	new_slice_5 := []string{"his", "name", "is", "john"}
	for i := 0; i < len(new_slice_5); i++ {
		fmt.Println(new_slice_5[i])
	}
	//iteration over range in for loop
	for ele, ok := range new_slice_5 {
		fmt.Printf("the index is : %d \n", ele)
		fmt.Printf("the value is : %v \n", ok)
	}
	//for only values we have to use blank identifier
	for _, ok := range new_slice_5 {
		fmt.Printf("the value is : %v \n", ok)
	}

}
