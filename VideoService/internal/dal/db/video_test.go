package db

import (
	"testing"
	"unsafe"
)

func TestArrayTest(t *testing.T) {
	//TODO 压缩算法
	var arr []int64
	for i := 0; i < 1000000; i++ {
		arr = append(arr, int64(i))
	}
	size := unsafe.Sizeof(arr)
	println(size)
}
