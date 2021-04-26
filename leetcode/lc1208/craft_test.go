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
	s1  interface{}
	s2  interface{}
	maxCost interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"acdbccd", "adcdcca", 2, 3},
		{"leetcode", "lceetdec", 3, 3},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := equalSubstring(_t.s1.(string), _t.s2.(string), _t.maxCost.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.s1.(string), _t.s2.(string), _t.maxCost.(int), _res, _t.exp)
		}
	}
}
