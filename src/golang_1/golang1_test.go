package golang_1

import (
	"fmt"
	"testing"
)

func TestSlice03(t *testing.T) {
	Slice03()
}

func TestAppend01(t *testing.T) {
	Append01()
}

func TestSliceCopy(t *testing.T) {
	SliceCopy()
	arr := []int{1, 3, 4}
	fmt.Println(arr[:0]) // arr[0:0] arr[:0] 意为清空切片
}

func TestMap(t *testing.T) {
	Map()
}

func TestSyncMap(t *testing.T) {
	SyncMap()
}

func TestList(t *testing.T) {
	List()
}
