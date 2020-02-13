// 這裡不會顯示在 godoc 建立的說明文件中
// 所以通常檔案的最上方會放置版權聲名及開發人員名單

// Package c03godoc 展示：在使用 godoc 來製作說明文件時，
// 註解一定要緊連 package、func 等關鍵字的「上方」才會顯示
//
// 在這邊的註解是緊連在【package godoc】上方才得以顯示
package c03godoc

import "fmt"

// Function01 是這份檔案的第 1 個函數
//
// 這個函數會印出「I am func Function01」文字
//
// 在這邊的註解是緊連在【func Function01() {】上方才得以顯示
func Function01() {
	fmt.Println("I am Function01")
}

/*
Function02 是這份檔案的第 2 個函數

這個函數會印出「I am func Function02」文字

在這邊的註解是緊連在【func Function02() {】上方才得以顯示
*/
func Function02() {
	fmt.Println("I am Function02")
}

// Function03 是這份檔案的第 3 個函數
/*
這個函數會印出「I am func Function03」文字
*/
// 在這邊的註解是緊連在【func Function03() {】上方才得以顯示
func Function03() {
	fmt.Println("I am Function03")
}

// function04 是這份檔案的第 4 個函數
//
// 這個函數會印出「I am func function04」文字
//
// 要注意的是，在這邊的註解雖然是緊連在【func Function04() {】上方，但因為此函數為私有函數，所以不會顯示出來
func function04() {
	fmt.Println("I am Function04")
}
