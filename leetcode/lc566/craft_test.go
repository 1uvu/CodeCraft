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
	r interface{}
	c interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[][]int{[]int{1,2}, []int{3,4}}, 1,4, [][]int{[]int{1,2,3,4}}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := matrixReshape(_t.in.([][]int), _t.r.(int), _t.c.(int)) // change there `in` type
		if !utils.CompareMatrix(_res, _t.exp.([][]int)) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
