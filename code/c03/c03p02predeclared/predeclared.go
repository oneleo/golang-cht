package main

import (
	"fmt"
)

// START OMIT
func main() {
	type int bool // HL
	i := [4]int{true, false, true, false}
	fmt.Printf("%10[1]T: %-15[2]p = %10.6[1]v\n", i[0], &i[0])
	fmt.Printf("%10[1]T: %-15[2]p = %10.6[1]v\n", i[1], &i[1])
	fmt.Printf("%10[1]T: %-15[2]p = %10.6[1]v\n", i[2], &i[2])
	fmt.Printf("%10[1]T: %-15[2]p = %10.6[1]v\n", i[3], &i[3])
}

// END OMIT
