package main

import (
	"fmt"

	"github.com/oneleo/golang-cht/code/c01/c01p01mymath"
	"github.com/yourname/hello/mymath"
)

func main() {
	a, b := "hello,", "world!"
	fmt.Println(a, b)
	c, d, e := c01p01mymath.AddMinusMultiply(1, 2)
	fmt.Println(c, d, e)
	f, g, h := mymath.AddMinusMultiply(4, 3)
	fmt.Println(f, g, h)
}
