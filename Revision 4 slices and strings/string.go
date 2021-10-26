//strings are immutable chains of arbitary bytes

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//creating a string
	a := "This is my string"
	fmt.Println(a)
	//backtics are the type of raw literals
	b := `I am a \n bad cop` //the escape sequence will not work on backtics
	fmt.Println(b)
	b1 := "I am a \n bad cop"
	fmt.Println(b1)
	//strings are immutable
	c := "foodie"
	//c[0] =="m" - it throws error
	fmt.Println(c)

	//looping through strings
	d := "I am Waiting"
	for ok, ele := range d {
		fmt.Printf("The index is %d and the value is %c\n", ok, ele)
	}
	//creating string from slice
	myslice1 := []byte{0x4, 0x65, 0x65, 0x6b, 0x73}
	mystring1 := string(myslice1)
	fmt.Println(mystring1)
	e := "hope is good"
	fmt.Println(len(e))            //length of the string
	f := utf8.RuneCountInString(e) //using this function
	fmt.Println(f)

}
