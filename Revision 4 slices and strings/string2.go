package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "!!how it is !!"
	b := strings.Trim(a, "!") //trimming
	fmt.Println(b)
	c := strings.TrimLeft(a, "!") //trimming at left
	fmt.Println(c)
	d := strings.TrimRight(a, "!") //trimming at right
	fmt.Println(d)
	f := "     this is a good     "
	e := strings.TrimSpace(f) //trimming the spaces
	fmt.Println(e)
	g := "HOW,ARE"
	h := strings.TrimSuffix(g, "ARE") //trimming the suffix
	fmt.Println(h)
	i := strings.TrimPrefix(g, "HOW") //trimming the prefix
	fmt.Println(i)
}
