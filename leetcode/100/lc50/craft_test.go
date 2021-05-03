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
	n   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{2.0, 10, 1024.0},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := myPow(_t.in.(float64), _t.n.(int)) // change there `in` type
		if _res != _t.exp.(float64) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
