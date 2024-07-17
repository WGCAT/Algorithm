func solution(survey []string, choices []int) string {
	category := [][]string{{"R", "T"}, {"C", "F"}, {"J", "M"}, {"A", "N"}}
	m := make(map[string]int)
	// for _, v := range category {
	// 	for _, vv := range v {
	// 		m[vv] = 0
	// 	}
	// }
	score := make([]int, len(choices))
	for i := 0; i < len(choices); i++ {
		score[i] = 4 - choices[i]
	}

	for i := 0; i < len(choices); i++ {
		if score[i] > 0 {
			str0 := string((survey[i])[0])
			m[str0] = m[str0] + score[i]
		} else if score[i] < 0 {
			str1 := string((survey[i])[1])
			m[str1] = m[str1] + (score[i] * (-1))
		}
	}

	var result string
	for _, v := range category {
		if m[v[0]] < m[v[1]] {
			result = result + v[1]
		} else {
			result = result + v[0]
		}
	}

	return result
}
