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
		{2, true},
		{3, false},
		{10, true},
		{27, false},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := judgeSquareSum(_t.in.(int)) // change there `in` type
		if _res != _t.exp.(bool) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
