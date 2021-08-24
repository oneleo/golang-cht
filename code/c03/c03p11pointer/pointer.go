package main

import "fmt"

func main() {
	// START OMIT
	v1, v2 := int8(1), int8(2)
	var ( // & 為取址運算子
		p1 *int8 = &v1
		p2 *int8 = &v2
	)
	fmt.Printf("1、存放 v1（= %v）的位址為：%v\n存放 v2（= %v）的位址為：%v\n\n", v1, p1, v2, p2)
	// * 為取值運算子
	fmt.Printf("2、回推 p1 位址對映的值為：%v\n回推 p2 位址對映的值為：%v\n\n", *p1, *p2)
	// 若嘗試修改 *p1、*p2 的值
	*p1, *p2 = 10, 20
	fmt.Printf("3、修改 *p1（to %v）及 *p2（to %v）後，\n", *p1, *p2)
	fmt.Printf("v1 現在的值為：%v\tv2 現在的值為：%v\n", v1, v2)
	// END OMIT
}
