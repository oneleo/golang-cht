// START OMIT
// ---------- （一）套件字句 ---------- //
// ↓ 這一行說該份檔案中所有的程式碼都歸屬於「main」套件。
package main

// ---------- （二）匯入段落 ---------- //
import (
	// ↓ 這裡說我們即將從「fmt」套件使用可讓文字格式化的程式碼。
	"fmt"
)

// ---------- （三）程式碼區 ---------- //
// ↓ 「main」是特別的函式；當您的程式啟動時它會第一個被執行。
func main() {
	fmt.Println("Hello, playground")
	// ↑ 這裡會調用「fmt」套件的函式「Println」；
	// 會在您的終端機（或者網頁瀏覽器，假如您是用 Go 遊樂場）顯示印出。
}

// END OMIT

// func Println(s string) {
// 	fmt.Println(s)
// }
