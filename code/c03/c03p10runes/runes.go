package main

import "fmt"

func main() {
	// START OMIT
	fmt.Printf("%02x\n", ' ')  // 20
	fmt.Printf("%02x\n", 'A')  // 41
	fmt.Printf("%02x\n", 'a')  // 61
	fmt.Printf("%02x\n", '\n') // 0A
	fmt.Printf("%02x\n", '\t') // 09
	fmt.Printf("%02x\n", '\\') // 5C
	// END OMIT
}
