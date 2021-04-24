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
	target interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1,2,3}, 4, 7},
		{[]int{1,2,4,5}, 9, 118},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := combinationSum4(_t.in.([]int), _t.target.(int))  // change there `in` type
		if _res != _t.exp.(int) { // change there `exp` type
			t.Error(_t.in, _t.target, _res, _t.exp)
		}
	}
}
