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
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := countBits(_t.in.(int)) // change there `in` type
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
