/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"

	"github.com/1uvu/codecraft/structures"
	"github.com/1uvu/codecraft/utils"
)

type SingleTest struct {
	in  interface{}
	exp interface{}
}

var null = utils.NULL

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1, null, 2, null, 5, 4}, []int{1, null, 2, null, 4, null, 5}},
		{[]int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9}, []int{1, null, 2, null, 3, null, 4, null, 5, null, 6, null, 7, null, 8, null, 9}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := increasingBST(structures.ArrayConvertToBinaryTree(_t.in.([]int))) // change there `in` type
		if !structures.IsBinaryTreeEqual(_res, structures.ArrayConvertToBinaryTree(_t.exp.([]int))) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
