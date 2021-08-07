package main

// START OMIT
const ( // 全域常數
	mp float64 = 1.67262e-27 // 質子質量
	me         = 9.10938e-31 // 電子質量
)
const Pi float32 = 3.14159 // 圓周率
var (                      // 全域字串變數
	name string = "gopher"
	age         = 18
)
var id string = "id0001"

func main() {
	const ( // 區域常數
		g float64 = 6.67384e-11 // 萬有引力常數
		c         = 2.99792e+10 // 真空中光速
	)
	var ( // 區域字串變數
		gender   string = "female"
		postcode        = 220223
	)
	// END OMIT
	_ = g
	_ = c
	_ = gender
	_ = postcode
}
