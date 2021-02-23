package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "I'm \"Gopher\"!"
	for k, v := range s {
		fmt.Printf("k：%d,v：%c == %d\n", k, v, v)
	}

	i := 10
	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)
	s3 := fmt.Sprintf("%d", i)
	fmt.Printf("%v, %v, %v\n", s1, s2, s3)

	w1 := "Aikido"
	fmt.Printf(w1 + "\n")
	fmt.Printf("Number of bytes %d\n", len(w1))
	fmt.Printf("Number of runes %d\n", utf8.RuneCountInString(w1))
	w2 := "合氣道"
	fmt.Printf(w2 + "\n")
	fmt.Printf("Number of bytes %d\n", len(w2))
	fmt.Printf("Number of runes %d\n", utf8.RuneCountInString(w2))

	//Case 1 s contains sep. Will output slice of length 3
	res := strings.Split("ab$cd$ef", "$")
	fmt.Println(res)
	//Case 2 s doesn't contain sep. Will output slice of length 1
	res = strings.Split("ab$cd$ef", "-")
	fmt.Println(res)
	//Case 3 sep is empty. Will output slice of length 8
	res = strings.Split("ab$cd$ef", "")
	fmt.Println(res)
	//Case 4 both s and sep are empty. Will output empty slice
	res = strings.Split("", "")
	fmt.Println(res)
	res2 := strings.ReplaceAll("abcdabxyabr", "a", "")
	fmt.Println(res2)

	value := "cat dog"
	// Get index of this substring.
	result := strings.Index(value, "dog")
	fmt.Println(result)

	input := "one two one"
	// Get last index, searching from right to left.
	ii := strings.LastIndex(input, "one")
	fmt.Println(ii)

	value1 := "rainbow"
	value2 := "the rainbow"
	// See if value1 is found in value2.
	if strings.Index(value2, value1) != -1 {
		fmt.Println("FOUND")
	}
}
