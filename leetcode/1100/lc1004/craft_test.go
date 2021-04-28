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
	K   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := longestOnes(_t.in.([]int), _t.K.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
