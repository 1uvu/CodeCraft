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

const null = utils.NULL
const False = utils.False
const True = utils.True

type SingleTest struct {
	opt interface{}
	in  interface{}
	exp interface{}
}

func Test(t *testing.T) { // rename function
	tests := []SingleTest{
		{[]string{"MyQueue", "push", "push", "peek", "pop", "empty"}, [][]int{{}, {1}, {2}, {}, {}, {}}, []int{null, null, null, 1, 1, False}},
		// ...
	}
	fmt.Println("begin testing...")
	for _, _t := range tests {
		queue := Constructor()
		_res := make([]int, len(_t.exp.([]int)))
		for i, opt := range _t.opt.([]string) {
			in := _t.in.([][]int)[i]
			switch opt {
			case "push":
				queue.Push(in[0])
				_res[i] = null
			case "peek":
				_res[i] = queue.Peek()
			case "pop":
				_res[i] = queue.Pop()
			case "empty":
				if queue.Empty() {
					_res[i] = True
				} else {
					_res[i] = False
				}
			default:
				_res[i] = null
			}
		}
		if !utils.CompareArray(_res, _t.exp.([]int)) {
			t.Error(_t.opt, _t.in, _res, _t.exp)
		}
	}
}
