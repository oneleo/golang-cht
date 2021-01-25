package main // OMIT START

/*
#include <stdio.h>

void pyramid(int p) {
    for(int i = 0 ; i <= p ; i++ ) {
        for(int j = 1 ; j <= i ; j++ ) {
        printf("*");
        }
    printf("\n");
    }
}

int divide(int a, int b){
    return a / b;
}
*/
import "C" // 注意：此語句上方不能有其他空白列

import (
	"fmt"
)

// Divide is ...
func Divide(a, b int) (out int) {
	return int(C.divide((C.int)(a), (C.int)(b)))
}

func main() {
	for i := 1; i <= 3; i++ {
		C.pyramid((C.int)(i)) // OMIT END
	}
	fmt.Println(Divide(6, 3))
}
