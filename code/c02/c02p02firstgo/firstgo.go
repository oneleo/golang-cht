// START OMIT
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, playground")
	Println("Hello, playground")
}

func Println(s string) {
	log.Println("My Println:", s)
}

// END OMIT
