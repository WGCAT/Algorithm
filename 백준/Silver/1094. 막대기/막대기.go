package main

import "fmt"

func main() {
	var X int
	fmt.Scanln(&X)

	barLen := 64
	sumLen := barLen
	var haveBar []int

	for i := 0; i < 7; i++ {
		// 막대 합계가 지민이 원하는 길이보다 크다면
		if sumLen > X {
			if sumLen == barLen {
				barLen = barLen / 2
				sumLen = barLen
				continue
			} else {
				sumLen = sumLen - barLen
			}

			// 막대 합계가 지민이 원하는 길이보다 작다면
		} else if sumLen < X {
			// 슬라이스에 자른 막대 담기
			haveBar = append(haveBar, barLen)

			// 막대 합계가 지민이 원하는 길이와 같다면
		} else if sumLen == X {
			// 슬라이스에 자른막대 넣고 for문 빠져나오기
			haveBar = append(haveBar, barLen)
			break
		}
		barLen = barLen / 2
		sumLen = sumLen + barLen

	}
	countBar := len(haveBar)
	fmt.Println(countBar)

}
