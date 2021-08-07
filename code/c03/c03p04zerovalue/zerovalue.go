package main

import "fmt"

// START OMIT
func main() {

	var q1 int            // 整數類型：0
	var q2 float32        // 浮點數類型：0.0
	var q3 bool           // 布林代數類型：false
	var q4 string         // 字串類型：""
	var q5 [1]int         // 陣列類型：[對映類型的零值]
	var q6 []int          // 切片類型：nil
	var q7 *int           // 指標類型：nil
	var q8 map[int]string // 映射類型：nil
	var q9 func(int) int  // 函式類型：nil

	fmt.Printf("q1 = %v\t\tq2 = %v\t\tq3 = %v\n", q1, q2, q3)
	fmt.Printf("q4 = %v\t\tq5 = %v\t\tq6 = %v\n", q4, q5, q6)
	fmt.Printf("q7 = %v\tq8 = %v\tq9 = %v\n", q7, q8, q9)
}

// END OMIT
