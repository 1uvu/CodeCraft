/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"
)

type SingleTest struct {
	in  interface{}
	id  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]*Employee{{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}}}, 1, 11},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := getImportance(_t.in.([]*Employee), _t.id.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
