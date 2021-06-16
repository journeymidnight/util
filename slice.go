package util

func StingInSlice(s string, ss []string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func IntInSlice(n int, ns []int) bool {
	for _, x := range ns {
		if x == n {
			return true
		}
	}
	return false
}

func MergeIntSlices(a, b []int) []int {
	m := make(map[int]bool)
	for _, x := range a {
		m[x] = true
	}
	for _, y := range b {
		m[y] = true
	}
	var result []int
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Keys(m map[int][]string) []int {
	result := make([]int, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
