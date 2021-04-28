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

const null int = utils.NULL

type SingleTest struct {
	init interface{}
	in   interface{}
	exp  interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{-2, 0, 3, -5, 2, -1}, [][2]int{{0, 2}, {2, 5}, {0, 5}}, []int{null, 1, -1, -3}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		obj := Constructor(_t.init.([]int))
		_res := make([]int, len(_t.exp.([]int)))
		_res[0] = null
		for i, in := range _t.in.([][2]int) {
			_res[i+1] = obj.SumRange(in[0], in[1])
		}
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.init, _t.in, _res, _t.exp)
		}
	}
}
