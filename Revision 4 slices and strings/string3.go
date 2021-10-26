package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "welcome, floks ,lets, rock"
	b := strings.Split(str, ",")
	fmt.Println(b)
	c := strings.SplitAfter(str, ",")
	fmt.Println(c)
	d := strings.SplitAfterN(str, "", 3)
	fmt.Println(d)
	//comparing the strings
	e := "STRING"
	f := "STRING"
	g := "string"
	//using assignment operators
	fmt.Println(e == f)
	fmt.Println(e != f)
	fmt.Println(e > g)
	fmt.Println(e >= f)
	//we can also use compare method
	fmt.Println(strings.Compare(e, g)) //returns -1 because e>g

}
