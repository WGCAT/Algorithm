package main

import (
	"fmt"
)

func main() {
	// 베라의 패션
	// 상의 N벌과 하의 N벌
	var N int
	fmt.Scanln(&N)

	// 카운트 선언 및 초기화
	count := 0

	// 첫번째 for문으로 상의 N벌 돌림
	for i := 0; i < N; i++ {
		// 두번째 for문으로 하의 N벌 돌림
		for j := 0; j < N; j++ {
			// 상의와 하의가 같은번째가 아니라면
			if i != j {
				// 카운트
				count++
			}
		}
	}

	fmt.Println(count)
}
