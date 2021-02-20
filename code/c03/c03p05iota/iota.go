package main

import (
	"fmt"
)

// START OMIT
const (
	Zero = iota // 0
	One         // 1
	Two         // 2
	_           // (skip)
	Four        // 4
)
const (
	Readable   uint8 = 1 << iota // 0001
	Writable                     // 0010
	Executable                   // 0100
)

func main() {
	fmt.Println("One =", One, "\tTwo =", Two, "\tFour =", Four)
	var police uint8 = 0b0111 // HL
	fmt.Printf("File police = %.4b:\n", police)
	police = police &^ Writable // 符合則清除 // HL
	fmt.Printf("Readable\t\t?\t%t\n", police&Readable == Readable)
	fmt.Printf("Writable\t\t?\t%t\n", police&Writable == Writable)
	fmt.Printf("Executable\t?\t%t\n", police&Executable == Executable)
}

// END OMIT
