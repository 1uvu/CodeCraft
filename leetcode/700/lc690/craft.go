/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/employee-importance/"
@Tags    :   [Map，DFS]
---------------------------
@Idea:
 - 使用 Map 存储 ID
 - DFS 遍历
*/
package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) (sum int) {
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}

	var DFS func(int)
	DFS = func(id int) {
		e := mp[id]
		sum += e.Importance
		for _, subId := range e.Subordinates {
			DFS(subId)
		}
	}
	DFS(id)
	return
}
