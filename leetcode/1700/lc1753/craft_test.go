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
	a   interface{}
	b   interface{}
	c   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{4, 4, 6, 7},
		{1, 8, 8, 8},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := maximumScore(_t.a.(int), _t.b.(int), _t.c.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.a.(int), _t.b.(int), _t.c.(int), _res, _t.exp)
		}
	}
}
