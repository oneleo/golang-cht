package main

import (
	"fmt"
	"strconv"
)

// START OMIT
func main() {
	i, j := 10, "10"
	str1 := strconv.Itoa(i) // 數字轉字串
	str2 := strconv.FormatInt(int64(i), 10)
	str3 := fmt.Sprintf("%d", i)
	fmt.Printf("數字轉字串：（str1）\"%v\",（str2）\"%v\",（str3）\"%v\"\n", str1, str2, str3)
	if i, err := strconv.Atoi(j); err == nil {
		fmt.Printf("字串轉數字：（int1）%v, ", i)
	}
	if i, err := strconv.ParseInt(j, 10, 64); err == nil {
		fmt.Printf("（int2）%v, ", i)
	}
	if _, err := fmt.Sscanf(j, "%d", &i); err == nil {
		fmt.Printf("（int3）%v\n", i)
	}
}

// END OMIT
