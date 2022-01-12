package t_test

import (
	"fmt"
	"testing"
)

func BenchmarkMapAsSetLarge(b *testing.B) {
	m := map[string]struct{}{}
	var temp bool
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		m[str] = struct{}{}
		_, ok := m[str]
		temp = ok
	}
	_ = temp
}

func BenchmarkSliceAsSetLarge(b *testing.B) {
	var m []string
	var temp bool
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		m = append(m, str)
		for i2 := range m {
			if m[i2] == str {
				temp = true
				break
			}
		}
	}
	_ = temp
}

func BenchmarkMapAsSetSmall(b *testing.B) {
	m := map[string]struct{}{}
	var temp bool
	count := 10
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		if count > 0 {
			count--
			m[str] = struct{}{}
			_, ok := m[str]
			temp = ok
		}
	}
	_ = temp
}

func BenchmarkSliceAsSetSmall(b *testing.B) {
	var m []string
	var temp bool
	count := 10
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("%d", i)
		m = append(m, str)
		if count > 0 {
			count--
			for i2 := range m {
				if m[i2] == str {
					temp = true
					break
				}
			}

		}
	}
	_ = temp
}
