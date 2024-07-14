func solution(name []string, yearning []int, photo [][]string) []int {
	m := make(map[string]int)
	var sum int
	result := []int{}
	for i := 0; i < len(name); i++ {
		m[name[i]] = yearning[i]
	}
	for _, p1 := range photo {
		sum = 0
		for _, p2 := range p1 {
			value, exsists := m[p2]
			if exsists {
				sum = sum + value
			}
		}
		result = append(result, sum)
	}
	return result
}