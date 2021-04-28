/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"

	"github.com/1uvu/codecraft/utils"
)

type SingleTest struct {
	A  interface{}
	B interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1,1,3}, []int{2,2,1}, []int{1,1}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := fairCandySwap(_t.A.([]int), _t.B.([]int)) // change there `in` type
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.A, _t.B, _res, _t.exp)
		}
	}
}
