/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"
)

type SingleTest struct {
	in  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}}, true},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := isToeplitzMatrix(_t.in.([][]int)) // change there `in` type
		if _res != _t.exp.(bool) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
