import (
	"strings"
)
func solution(survey []string, choices []int) string {
	category := [][]string{{"R", "T"}, {"C", "F"}, {"J", "M"}, {"A", "N"}}
	m := make(map[string]int)
	for _, v := range category {
		for _, vv := range v {
			m[vv] = 0
		}
	}
	score := make([]int, len(choices))
	for i := 0; i < len(choices); i++ {
		score[i] = 4 - choices[i]
	}

	for i := 0; i < len(choices); i++ {
		one := strings.Split(survey[i], "")
		if score[i] > 0 {
			m[one[0]] = m[one[0]] + score[i]
		} else if score[i] < 0 {
			m[one[1]] = m[one[1]] + (score[i] * (-1))
		}
	}

	var result []string
	for _, v := range category {
		if m[v[0]] < m[v[1]] {
			result = append(result, v[1])
		} else {
			result = append(result, v[0])
		}
	}

	return strings.Join(result, "")
}
