package t_test

import (
	"fmt"
	"testing"
)

func BenchmarkMapAsSetDelete(b *testing.B) {
	m := map[string]struct{}{}
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		m[str] = struct{}{}
		delete(m, str)
	}
}

func BenchmarkMapAsSetBool(b *testing.B) {
	m := map[string]bool{}
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		m[str] = true
		m[str] = false
	}
}
