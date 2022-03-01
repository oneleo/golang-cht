package main

import (
	"fmt"

	"github.com/oneleo/golang-cht/code/c01/c01p01mymath"
	"github.com/yourname/hello/mymath"
	"github.com/yourname/hello/mypack"
)

func main() {
	a, b := "hello,", "world!"
	fmt.Println(a, b)
	c, d, e := c01p01mymath.AddMinusMultiply(1, 2)
	fmt.Println(c, d, e)
	f, g, h := mymath.AddMinusMultiply(4, 3)
	fmt.Println(f, g, h)
	fmt.Println("mypack.var1 =", mypack.Var1) // private variable
	fmt.Println("mypack.Var2 =", mypack.Var2)
	mypack.Myfunc1(4) // private function
	mypack.Myfunc2(3)
}
