// 메뉴에 아메리카노 포함 4500, 카페라테가 포함 5000, else 4500

import "strings"

func solution(order []string) int {
	// 금액을 담을 변수 선언
    var total int
    // for range문을 사용하여 order배열의 값을 차례로 menu에 담아 돌린다
	for _, menu := range order {
        // menu에 americano가 포함되어 있을 경우 4500
		if strings.Contains(menu, "americano") {
			total = total + 4500
        // menu에 cafelatte가 포함되어 있을 경우 5000
		} else if strings.Contains(menu, "cafelatte") {
			total = total + 5000
        // 그 외의 경우 4500
		} else {
			total = total + 4500
		}

	}
	return total
}