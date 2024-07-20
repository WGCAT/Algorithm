func solution(clothes [][]string) int {
	clotgesMap := make(map[string]int)
	result := 1
	for _, v := range clothes {
		clotgesMap[v[1]]++
	}
	for _, v := range clotgesMap {
		result = (result) * (v + 1)
	}
	return result - 1
}
