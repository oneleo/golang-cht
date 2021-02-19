package main

import (
	"fmt"
)

// START OMIT
func main() {
	const feet float64 = 3.2808399                           // HL
	quantity, length, width, customerName := declaration03() // HL
	fmt.Println(customerName)
	fmt.Println("買了", quantity, "雙鞋子，鞋子尺寸為", length*feet*width*feet, "平方英呎")
}
func declaration01() (int, float64, float64, string) {
	var quantity int          // 宣告變數：「var」鍵詞 + 變數名稱 + 型別
	var length, width float64 // 您可以同時宣告一樣型別的多個變數
	quantity = 4              // 指派變數
	length, width = 1.2, 2.4  // 同時指派多個變數 // HL
	return quantity, length, width, "Damon Cole's sister"
}
func declaration02() (int, float64, float64, string) {
	var quantity int = 4         // 宣告變數以及指派變數
	var length, width = 1.2, 2.4 // 型別推斷 // HL
	return quantity, length, width, "Damon Cole's brother"
}
func declaration03() (int, float64, float64, string) {
	quantity := 4 // 短變數宣告
	length, width := 1.2, 2.4
	quantity, length, width, customerName := 4, 1.2, 2.4, "Damon Cole" // 至少要有一個新變數 // HL
	return quantity, length, width, customerName
}

// END OMIT
