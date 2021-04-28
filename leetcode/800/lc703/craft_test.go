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
	k    interface{}
	init interface{}
	add  interface{}
	exp  interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{3, []int{4, 5, 8, 2}, []int{3, 5, 10, 9, 4}, []int{null, 4, 5, 5, 8, 8}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := []int{null}
		kth := Constructor(_t.k.(int), _t.init.([]int))
		for _, val := range _t.add.([]int) {
			_res = append(_res, kth.Add(val))
		}
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.k, _t.init, _t.add, _res, _t.exp)
		}
	}
}
