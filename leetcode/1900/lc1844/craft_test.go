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
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"a1c1e1", "abcdef"},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := replaceDigits(_t.in.(string)) // change there `in` type
		if _res != _t.exp.(string) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
