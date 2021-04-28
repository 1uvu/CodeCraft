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
	in  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{42}, []int{0}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := dailyTemperatures(_t.in.([]int)) // change there `in` type
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
