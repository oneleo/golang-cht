package main

import "fmt"

// START OMIT
type Cart struct {
	ID   string
	Paid bool
}

func main() {
	cart := Cart{
		ID:   "121514",
		Paid: true,
	}
	cartPtr := &cart
	cartDeref := *cartPtr
	fmt.Printf("cart 的位址為：%v\t回推 cartPtr 位址的值為：%v", cartPtr, cartDeref)
}

// END OMIT
