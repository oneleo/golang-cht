// 這裡不會顯示在 godoc 建立的說明文件中
// 所以通常檔案的最上方會放置版權聲名及開發人員名單

// Package c01 是在中華電信學院開設的 Go 語言課程第一章節所建立的套件
package c01

// MultiMath 是一個雙重計算功能的函式,
// 此函式您需要輸入 2 個整數：a 和 b.
// 在您呼叫函式，執行後會有 2 筆輸出:
// 第 1 個輸出是加法結果：a + b.
// 第 2 個輸出是減法結果：a - b.
func MultiMath(a, b int) (add, minus int) {
	return a + b, a - b
}

// https://pkg.go.dev/github.com/oneleo/golang-cht
