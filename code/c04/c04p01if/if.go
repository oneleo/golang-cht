package main

import "fmt"

func main() {
	// START OMIT
	a, b, c := 1, 2, 3
	if a < b {
		fmt.Printf("1、a（%v）小於 b（%v）\n", a, b)
	}
	if a < b && b < c {
		fmt.Printf("2、a（%v）小於 b（%v），且 b（%v）小於 c（%v）\n", a, b, b, c)
	}
	if a := 99; a < b {
		fmt.Printf("3、a（%v）小於 b（%v）\n", a, b)
	} else {
		fmt.Printf("3、a（%v）大於 b（%v）\n", a, b)
	}
	fmt.Println("4、a 現在的值到底是？（%v）", a)
	// END OMIT
}
