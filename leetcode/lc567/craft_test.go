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
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"ab", "eidbaooo",true},
		{"ect","leetcode",true},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := checkInclusion(_t.s1.(string), _t.s2.(string)) // change there `in` type
		if _res != _t.exp.(bool) {
			t.Error(_t.s1.(string), _t.s2.(string), _res, _t.exp)
		}
	}
}
