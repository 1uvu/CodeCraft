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

const null int = utils.NULL

type SingleTest struct {
	in   interface{}
	low  interface{}
	high interface{}
	exp  interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{10, 5, 15, 3, 7, null, 18}, 7, 15, 32},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := rangeSumBST(structures.ArrayConvertToBinaryTree(_t.in.([]int)), _t.low.(int), _t.high.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
