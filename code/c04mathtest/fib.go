package c04mathtest

func Fib(n int) int { // OMIT START
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2) // OMIT END
}
