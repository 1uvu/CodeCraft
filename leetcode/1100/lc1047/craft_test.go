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
		{"abbaca", "ca"},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := removeDuplicates(_t.in.(string)) // change there `in` type
		fmt.Println(_res)
		if _res != _t.exp.(string) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
