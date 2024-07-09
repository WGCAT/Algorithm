func solution(topping []int) int {
    cs := make([]int, 1000000)
	bro := make([]int, 1000000)
	csCount := 0
	broCount := 0
	result := 0
	for _, v := range topping {
		if cs[v] == 0 {
			csCount++
		}
		cs[v]++
	}
	for _, v := range topping {
		if bro[v] == 0 {
			broCount++
		}
		bro[v]++
		cs[v]--
		if cs[v] == 0 {
			csCount--
		}
		if csCount == broCount {
			result++
		}
	}
	return result
}
