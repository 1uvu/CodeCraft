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

var null int = utils.NULL

type SingleTest struct {
	in  interface{}
	k   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{3, 1, 4, null, 2, null, 7}, 4, 4},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := kthSmallest(structures.ArrayConvertToBinaryTree(_t.in.([]int)), _t.k.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
