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
		{"[]{}()", true},
		{"[(){}[]", false},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := isValid(_t.in.(string)) // change there `in` type
		if _res != _t.exp.(bool) {      // change there `exp` type
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
