package t

import (
	"fmt"
	"os"
)

func MakeStringShort() (v1 string, v2 string) {
	pid := os.Getpid()
	v1 = fmt.Sprintf("short %d", pid)
	v2 = fmt.Sprintf("short %d", pid)
	return
}

func MakeStringLong() (v1 string, v2 string) {
	pid := os.Getpid()
	for i := 0; i < 10000; i++ {
		v1 += fmt.Sprintf("%d", pid)
		v2 += fmt.Sprintf("%d", pid)
	}
	return
}
