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
	D   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1,2,3,4,5,6,7,8,9,10}, 5, 15},
		{[]int{3,2,2,4,1,4}, 3, 6},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := shipWithinDays(_t.in.([]int), _t.D.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _t.D, _res, _t.exp)
		}
	}
}
