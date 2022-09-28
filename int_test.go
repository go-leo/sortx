package sortx

import (
	"testing"
)

func TestIntsDesc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	IntsDesc(arr)
	t.Log(arr)
}

func TestIntsAsc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	IntsAsc(arr)
	t.Log(arr)
}
