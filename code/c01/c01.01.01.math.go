// 這裡不會顯示在 godoc 建立的說明文件中
// 所以通常檔案的最上方會放置版權聲名及開發人員名單

// Package c01 是在中華電信學院開設的 Go 語言課程第一章節所建立的套件，
/*
此套件包含了 2 個函式，

一個是有複合計算功能的函式 AddMinusMultiply，
*/
// 另一個是加總計算功能的函式 Sum。
package c01

import "fmt"

/*
AddMinusMultiply 是一個三重計算功能的函式，

此函式您需要輸入 2 個整數：a 和 b，

當您呼叫函式，執行後將會有 3 筆輸出：

第 1 個輸出是加法結果：a + b

第 2 個輸出是減法結果：a - b

第 3 個輸出是乘法結果：a * b
*/
func AddMinusMultiply(a, b int) (add, minus, multiply int) {
	return a + b, a - b, a * b
}

// Sum 是一個加總計算功能的函式，
//
// 此函式您可以輸入任意數量的整數，
//
// 當您呼叫函式，執行後會將輸入的整數加總後輸出。
func Sum(in ...int) (sum int) {
	s := 0
	for _, i := range in {
		s += i
	}
	return s
}

// privateFunc 是私有函式，
//
// 這個函式會印出「I'm private function.」文字，
//
// 但要注意的是，因為此函式為私有函式，所以不會在說明文件中顯示出來。
func privateFunc() {
	fmt.Println("I'm private function.")
}

// 註：觀看文件的方式：
// 1、godoc -http=:6060 -play
// 2、https://pkg.go.dev/github.com/oneleo/golang-cht
