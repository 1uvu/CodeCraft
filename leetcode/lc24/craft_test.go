/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"

	"github.com/1uvu/codecraft/structures"
)

type SingleTest struct {
	in  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{3, 2, 4, 5, 8}, []int{2, 3, 5, 4, 8}},
		{[]int{1,7,5,2,4,1,3}, []int{7,1,2,5,1,4,3}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := swapPairs_(structures.ArrayConvertToLinkedList(_t.in.([]int)))
		if !structures.IsLinkedListEqual(_res, structures.ArrayConvertToLinkedList(_t.exp.([]int))) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
