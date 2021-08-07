package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// START OMIT
func main() {
	fmt.Printf("1-1 評測 FmtSprintf\t%v\n", testing.Benchmark(BenchmarkFmtSprintf))
	fmt.Printf("1-2 評測 StrAdd\t%v\n", testing.Benchmark(BenchmarkStrAdd))
	fmt.Printf("1-3評測 StrJoin\t%v\n", testing.Benchmark(BenchmarkStrJoin))
	fmt.Printf("1-4 評測 StrBuffer\t%v\n", testing.Benchmark(BenchmarkStrBuffer))

	fmt.Printf("\n2-1 評測 FmtSprintfOne\t%v\n", testing.Benchmark(BenchmarkFmtSprintfOne))
	fmt.Printf("2-2 評測 StrAddOne\t%v\n", testing.Benchmark(BenchmarkStrAddOne))
	fmt.Printf("2-3 評測 StrJoinOne\t%v\n", testing.Benchmark(BenchmarkStrJoinOne))
	fmt.Printf("2-4 評測 StrBufferOne\t%v\n", testing.Benchmark(BenchmarkStrBufferOne))
}

// END OMIT

// BenchmarkFmtSprintf 使用 fmt.Sprintf 進行接續串接
func BenchmarkFmtSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s%s%s", s, "hello", "world")
	}
}

// BenchmarkStrAdd 使用 + 運算進行接續串接
func BenchmarkStrAdd(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "hello" + "world"
	}
}

// BenchmarkStrJoin 使用 strings.Join 進行接續串接
func BenchmarkStrJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = strings.Join([]string{s, "hello", "world"}, "")
	}
}

// BenchmarkStrBuffer 使用 bytes.Buffer 進行接續串接
func BenchmarkStrBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("hello")
		buffer.WriteString("world")
	}
}

// BenchmarkFmtSprintfOne 使用 fmt.Sprintf 進行單一串接
func BenchmarkFmtSprintfOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s", "hello", "world")
		_ = s
	}
}

// BenchmarkStrAddOne 使用 + 運算進行單一串接
func BenchmarkStrAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "hello" + "world"
		_ = s
	}
}

// BenchmarkStrJoinOne 使用 strings.Join 進行單一串接
func BenchmarkStrJoinOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join([]string{"hello", "world"}, "")
		_ = s
	}
}

// BenchmarkStrBufferOne 使用 bytes.Buffer 進行單一串接
func BenchmarkStrBufferOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.WriteString("hello")
		b.WriteString("world")
	}
}
