package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

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
	fmt.Printf("拆解：%v\n取代：%v\n從左向右查詢：%v\n從右向左查詢：%v\n查詢失敗：%v\n", res1, res2, res3, res4, res5)

	i, j := 10, "10"
	str1 := strconv.Itoa(i) // 數字轉字串
	str2 := strconv.FormatInt(int64(i), 10)
	str3 := fmt.Sprintf("%d", i)
	fmt.Printf("數字轉字串：\"%v\", \"%v\", \"%v\"\n", str1, str2, str3)
	if i, err := strconv.Atoi(j); err == nil {
		fmt.Printf("字串轉數字：%v, ", i)
	}
	if i, err := strconv.ParseInt(j, 10, 64); err == nil {
		fmt.Printf("%v, ", i)
	}
	if _, err := fmt.Sscanf(j, "%d", &i); err == nil {
		fmt.Printf("%v\n", i)
	}
}
