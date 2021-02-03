package main

import (
	"fmt"
)

// START OMIT
/*
#include <stdio.h>
int divide(int a, int b){
    return a / b;
}
*/
import "C" // 注意：此語句上方不能有其他空白列	// HL

// Divide 可進行兩個整數的除法。
func Divide(a, b int) (out int) {
	return int(C.divide((C.int)(a), (C.int)(b))) // HL
}

func main() {
	fmt.Println(Divide(6, 3))
}

// END OMIT
