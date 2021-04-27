/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/pascals-triangle-ii/"
@Tags    :   [动态规划，数学问题]
---------------------------
@Idea:
 - 利用数学规律，row[i] = row[i-1] * (rowIndex - i + 1) / i
 - 从前后两个方向计算第 k 行：row[rowIndex-i] = row[i]
*/
package leetcode

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	row[rowIndex] = 1
	for i:=1; i<rowIndex/2+1; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
		row[rowIndex-i] = row[i]
	}
	return row
}
