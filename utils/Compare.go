package utils

func CompareArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[0] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
