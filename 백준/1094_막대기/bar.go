package algorithm

func bar(jiminLen int) int {
	jiminLen = 15
	barLen := 64
	sumLen := barLen
	var haveBar []int

	for {
		// 막대 합계가 지민이 원하는 길이보다 크다면
		if sumLen > jiminLen {
			barLen = barLen / 2
			sumLen = barLen

			// 막대 합계가 지민이 원하는 길이보다 작다면
		} else if sumLen < jiminLen {
			// 슬라이스에 자른 막대 담기
			haveBar = append(haveBar, barLen)
			// 막대 합계에 자른막대 더하기
			sumLen = sumLen + barLen/2
			// 막대 자르기
			barLen = barLen / 2

			// 막대 합계가 지민이 원하는 길이와 같다면
		} else if sumLen == jiminLen {
			// 슬라이스에 자른막대 넣고 for문 빠져나오기
			haveBar = append(haveBar, barLen)
			break

		}
	}
	countBar := len(haveBar)
	return countBar

}
