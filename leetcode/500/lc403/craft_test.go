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
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1, 3, 4, 5, 7, 9, 10, 12}, true},
		{[]int{0, 1, 3, 6, 7}, false},
		{[]int{0, 2}, false},
		{[]int{0, 1, 3, 6, 10, 13, 15, 18}, true},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := canCross(_t.in.([]int)) // change there `in` type
		if _res != _t.exp.(bool) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
