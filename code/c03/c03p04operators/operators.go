package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Printf("true || false\t?\t%t\n", true || false)
	fmt.Printf("\"abc\" == \"abc\"\t?\t%t\n", "abc" == "abc")
	fmt.Printf("36 / 7\t\t=\t%d\n", 36/7)
	fmt.Printf("36 %% 7\t\t=\t%d\n", 36%7)
	fmt.Printf("0b1001 | 0b1011 \t=\t%.4[1]b\t=\t%[1]d\n", 0b1001|0b1011)
	fmt.Printf("0b1001 & 0b1011 \t=\t%.4[1]b\t=\t%[1]d\n", 0b1001&0b1011)
	fmt.Printf("0b1001 ^ 0b1011 \t=\t%.4[1]b\t=\t%[1]d\n", 0b1001^0b1011)
	fmt.Printf("0b1001 &^ 0b1011 \t=\t%.4[1]b\t=\t%[1]d\n", 0b1001&^0b1011)
	fmt.Printf("0b0100 << 1 \t=\t%.4[1]b\t=\t%[1]d\n", 0b0100<<1)
	fmt.Printf("0b0100 >> 1 \t=\t%.4[1]b\t=\t%[1]d\n", 0b0100>>1)
	i := 2
	i++ // ++ 與 -- 只能在一行陳述中置於變數後方，不能寫成 fmt.Println(i++)
	fmt.Printf("i++ →\ti\t=\t%d\n", i)
	i -= 2
	fmt.Printf("i-=2 →\ti\t=\t%d\n", i)
}

// END OMIT
