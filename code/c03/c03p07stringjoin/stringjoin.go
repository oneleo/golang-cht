package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func main() {
	fmt.Printf("BenchmarkFmtSprintf\t%v\n", testing.Benchmark(BenchmarkFmtSprintf))
	fmt.Printf("BenchmarkStrAdd\t\t%v\n", testing.Benchmark(BenchmarkStrAdd))
	fmt.Printf("BenchmarkStrJoin\t%v\n", testing.Benchmark(BenchmarkStrJoin))
	fmt.Printf("BenchmarkStrBuffer\t%v\n", testing.Benchmark(BenchmarkStrBuffer))

	fmt.Printf("BenchmarkFmtSprintfOne\t%v\n", testing.Benchmark(BenchmarkFmtSprintfOne))
	fmt.Printf("BenchmarkStrAddOne\t%v\n", testing.Benchmark(BenchmarkStrAddOne))
	fmt.Printf("BenchmarkStrJoinOne\t%v\n", testing.Benchmark(BenchmarkStrJoinOne))
	fmt.Printf("BenchmarkStrBufferOne\t%v\n", testing.Benchmark(BenchmarkStrBufferOne))
}

// fmt.Printf
func BenchmarkFmtSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += fmt.Sprintf("%s%s", "hello", "world")
	}
}

// 直接串接
func BenchmarkStrAdd(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "hello" + "world"
	}
}

// strings.Join
func BenchmarkStrJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strings.Join([]string{"hello", "world"}, "")
	}
}

// bytes.Buffer
func BenchmarkStrBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("hello")
		buffer.WriteString("world")
	}
}

func BenchmarkFmtSprintfOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s", "hello", "world")
		_ = s
	}
}

func BenchmarkStrAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "hello" + "world"
		_ = s
	}
}
func BenchmarkStrJoinOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join([]string{"hello", "world"}, "")
		_ = s
	}
}
func BenchmarkStrBufferOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.WriteString("hello")
		b.WriteString("world")
	}
}
