/* -*- encoding: utf-8 -*- */
/*
@File    :   craft_test.go
*/
package leetcode

import (
	"fmt"
	"testing"
)

type IntTest struct {
	in  int
	exp int
}

func TestFoo(t *testing.T) { // rename function
	tests := []IntTest{
		{1, 1},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := Foo(_t.in)
		if _res != _t.exp {
			t.Errorf("Fib(%d) = %d; expected %d", _t.in, _res, _t.exp)
		}
	}
}
