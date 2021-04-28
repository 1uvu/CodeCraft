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
		{[]int{1, 2, 4, 7, -1, 0}, true},
		{[]int{1, 2, 4, 3, 2, 5}, false},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := check(_t.in.([]int)) // change there `in` type
		if _res != _t.exp.(bool) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
