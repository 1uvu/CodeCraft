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
		{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
		{[]int{4, 8, 12, 16}, 2},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := maxTurbulenceSize(_t.in.([]int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
