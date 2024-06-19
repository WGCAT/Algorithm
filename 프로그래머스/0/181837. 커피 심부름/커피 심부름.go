// 메뉴에 아메리카노 포함 4500, 카페라테가 포함 5000, else 4500

import "strings"

func solution(order []string) int {
	var total int
	for _, menu := range order {
		if strings.Contains(menu, "americano") {
			total = total + 4500
		} else if strings.Contains(menu, "cafelatte") {
			total = total + 5000
		} else {
			total = total + 4500
		}

	}
	return total
}