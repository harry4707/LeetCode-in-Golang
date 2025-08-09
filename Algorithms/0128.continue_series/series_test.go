package _128_最长连续序列

import (
	"fmt"
	"testing"
)

func Test_continue_series(T *testing.T) {
	var array = []int{1, 2, 3, 8, 4, 7, 15, 17, 18, 19, 20, 56}
	result := ContinuousSequence(array)
	fmt.Println(result)
}
