package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var X int
	fmt.Scanln(&X)

	// X를 2진수인 string으로 변환
	strX := strconv.FormatInt(int64(X), 2)

	// 2진수 string을 슬라이스에 담기
	sliceX := strings.Split(strX, "")
	// 카운트 선언
	count := 0

	// for range문으로 2진수 슬라이스를 binary에 담아 돌림
	for _, binary := range sliceX {
		// binary가 1이라면
		if strings.Contains(binary, "1") {
			// 카운트 증가
			count++
		}
	}
	fmt.Println(count)

}