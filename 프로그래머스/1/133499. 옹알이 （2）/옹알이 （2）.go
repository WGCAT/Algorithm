import "strings"
func solution(babbling []string) int {
	say := [4]string{"aya", "ye", "woo", "ma"}
	var result int
	var choice []string
	var ru string
	for _, v := range babbling {
		for _, vv := range say {
			if strings.Contains(v, vv+vv) {
				goto nope
			}
		}
		choice = append(choice, v)
	nope:
	}
	for _, v := range choice {
		ru = v
		for _, vv := range say {
			ru = strings.Replace(ru, vv, "|", -1)
		}
		ru = strings.Replace(ru, "|", "", -1)
		if ru == "" {
			result++
		}
	}
	return result
}
