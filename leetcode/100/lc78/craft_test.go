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
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := subsets(_t.in.([]int)) // change there `in` type
		if !utils.CompareMatrix(_res, _t.exp.([][]int)) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
