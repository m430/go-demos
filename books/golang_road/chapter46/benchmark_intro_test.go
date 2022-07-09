package chapter46

import (
	"fmt"
	"strings"
	"testing"
)

var sl = []string{
	"Rob Pike",
	"Robert Griesemer",
	"Ken Thompson",
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByJoin(sl)
	}
}
