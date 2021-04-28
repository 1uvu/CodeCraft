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
	customers  interface{}
	grumpy  interface{}
	X  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1},3,16},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		_res := maxSatisfied(_t.customers.([]int),_t.grumpy.([]int),_t.X.(int)) // change there `in` type
		if _res != _t.exp.(int) {
			t.Error(_t.customers,_t.grumpy,_t.X, _res, _t.exp)
		}
	}
}
