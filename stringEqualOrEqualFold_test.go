package t_test

import (
	"strings"
	"testing"

	t "github.com/wirekang/go-benchmarks"
)

func BenchmarkStringContainsShort(b *testing.B) {
	v1, v2 := t.MakeStringShort()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.Contains(v1, v2)
	}
	_ = temp
}

func BenchmarkStringContainsLong(b *testing.B) {
	v1, v2 := t.MakeStringLong()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.Contains(v1, v2)
	}
	_ = temp
}

func BenchmarkStringEqualShort(b *testing.B) {
	v1, v2 := t.MakeStringShort()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = v1 == v2
	}
	_ = temp
}

func BenchmarkStringEqualLong(b *testing.B) {
	v1, v2 := t.MakeStringLong()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = v1 == v2
	}
	_ = temp
}

func BenchmarkStringEqualFoldShort(b *testing.B) {
	v1, v2 := t.MakeStringShort()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.EqualFold(v1, v2)
	}
	_ = temp
}

func BenchmarkStringEqualFoldLong(b *testing.B) {
	v1, v2 := t.MakeStringLong()
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.EqualFold(v1, v2)
	}
	_ = temp
}
