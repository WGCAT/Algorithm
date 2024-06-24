// A ~ Z : 65 ~ 90, a ~ z : 97 ~ 122

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)
	// 결과 담을 문자열 선언
	var result string

	// 주어진 문자열을 잘라 슬라이스에 담기
	slice := strings.Split(s1, "")

	// for range문으로 slice의 값을 v에 담아 돌림
	for _, v := range slice {
		// v가 a~z라면
		if v >= "a" && v <= "z" {
			// 대문자로 변환
			word := strings.ToUpper(v)
			result = result + word
			// 아닐경우
		} else {
			// 소문자로 변환
			word := strings.ToLower(v)
			result = result + word
		}
	}
	fmt.Println(result)
}
