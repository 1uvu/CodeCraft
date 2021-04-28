/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   ""
@Tags    :   []
---------------------------
@Idea:
 -
 -
*/
package leetcode

import "math"

func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		if isSquare(float64(c - a*a)) {
			return true
		}
	}
	return false
}

func isSquare(N float64) bool {
	n := math.Sqrt(N)
	return n == float64(int(n))
}
