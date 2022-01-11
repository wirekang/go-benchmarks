package stringcompareandcontains_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkStaticStringCompare(b *testing.B) {
	v1 := "asdfaqwer asdf"
	v2 := "asdfaqwer asdf"
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = v1 == v2
	}
	_ = temp
}

func BenchmarkStringCompare(b *testing.B) {
	pid := os.Getpid()
	v1 := fmt.Sprintf("short %d", pid)
	v2 := fmt.Sprintf("short %d", pid)
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = v1 == v2
	}
	_ = temp
}

func BenchmarkLongStringCompare(b *testing.B) {
	pid := os.Getpid()
	var v1, v2 string
	for i := 0; i < 10000; i++ {
		v1 += fmt.Sprintf("%d", pid)
		v2 += fmt.Sprintf("%d", pid)
	}

	var temp bool
	for i := 0; i < b.N; i++ {
		temp = v1 == v2
	}
	_ = temp
}

func BenchmarkStaticStringContains(b *testing.B) {
	v1 := "asdfaqwer asdf"
	v2 := "asdfaqwer asdf"
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.Contains(v1, v2)
	}
	_ = temp
}

func BenchmarkStringContains(b *testing.B) {
	pid := os.Getpid()
	v1 := fmt.Sprintf("short %d", pid)
	v2 := fmt.Sprintf("short %d", pid)
	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.Contains(v1, v2)
	}
	_ = temp
}

func BenchmarkLongStringContains(b *testing.B) {
	pid := os.Getpid()
	var v1, v2 string
	for i := 0; i < 10000; i++ {
		v1 += fmt.Sprintf("%d", pid)
		v2 += fmt.Sprintf("%d", pid)
	}

	var temp bool
	for i := 0; i < b.N; i++ {
		temp = strings.Contains(v1, v2)
	}
	_ = temp
}
