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
		{[]int{1,3,-1,-2,50,5}, 4, 13.0},
		{[]int{1,2,4,-2}, 2, 3.0},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := findMaxAverage(_t.in.([]int), _t.k.(int)) // change there `in` type
		if _res != _t.exp.(float64) {
			t.Error(_t.in, _t.k, _res, _t.exp)
		}
	}
}
