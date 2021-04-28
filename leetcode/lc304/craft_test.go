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
		{[][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}, [][4]int{{2, 1, 4, 3}, {1, 1, 2, 2}, {1, 2, 2, 4}}, []int{null, 8, 11, 12}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		obj := Constructor(_t.init.([][]int))
		_res := make([]int, len(_t.exp.([]int)))
		_res[0] = null
		for i, in := range _t.in.([][4]int) {
			_res[i+1] = obj.SumRegion(in[0], in[1], in[2], in[3])
		}
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.init, _t.in, _res, _t.exp)
		}
	}
}
