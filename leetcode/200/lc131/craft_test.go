/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"

	"github.com/1uvu/codecraft/utils"
)

type SingleTest struct {
	in  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := partition(_t.in.(string)) // change there `in` type
		if !utils.CompareStringMatrix(_res, _t.exp.([][]string)) {
			t.Error(_t.in, _res, _t.exp)
		}
	}
}
