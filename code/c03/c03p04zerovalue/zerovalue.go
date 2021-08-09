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

	fmt.Printf("q1=0? %v\tq2=0? %v\tq3=false? %v\n", q1 == 0, q2 == 0, q3 == false)
	fmt.Printf("q4=\"\"? %v\tq5=[0]? %v\tq6=nil? %v\n", q4 == "", q5 == [1]int{0}, q6 == nil)
	fmt.Printf("q7=nil? %v\tq8=nil? %v\tq9=nil? %v\n", q7 == nil, q8 == nil, q9 == nil)
}

// END OMIT
