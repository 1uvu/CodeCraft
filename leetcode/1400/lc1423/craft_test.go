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
	k   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := maxScore(_t.in.([]int), _t.k.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
