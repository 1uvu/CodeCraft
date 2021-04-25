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
	k   interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := characterReplacement(_t.in.(string), _t.k.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
