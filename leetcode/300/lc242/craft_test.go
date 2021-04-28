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
	s1 string
	s2 string
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"asdfgh", "adfgsh"},
		{"leetcode", "lcoteeed"},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := isAnagram(_t.s1, _t.s2) // change there `in` type
		if !_res {                      // change there `exp` type
			t.Error(_t.s1, _t.s2, _res)
		}
	}
}
