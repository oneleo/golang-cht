package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// START OMIT
func main() {
	w1, w2, w3 := "Hi,", "世界！❤", "a,p,p"
	for i, v := range w1 {
		fmt.Printf("i：%d, v：%c == 0x%X\n", i, v, v)
	}
	fmt.Printf("「%s」\t\t=> Byte 數：%d\t字數：%d\n", w1, len(w1), utf8.RuneCountInString(w1))
	fmt.Printf("「%s」\t=> Byte 數：%d\t字數：%d\n", w2, len(w2), utf8.RuneCountInString(w2))
	res1 := strings.Split(w3, ",")           // 拆解
	res2 := strings.ReplaceAll(w3, ",", "&") // 取代
	res3 := strings.Index(w3, "p")           // 從左向右查詢
	res4 := strings.LastIndex(w3, "p")       // 從右向左查詢
	res5 := strings.LastIndex(w3, "b")       // 查詢失敗（= -1）
	fmt.Printf("拆解：%v\n取代：%v\n從左向右查詢：%v\n", res1, res2, res3)
	fmt.Printf("從右向左查詢：%v\n查詢失敗：%v\n", res4, res5)
}

// END OMIT
