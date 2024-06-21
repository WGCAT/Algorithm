package main

import (
	"fmt"
	"strings"
)

func main() {
	// 단어 길이 재기
	var X string
	fmt.Scanln(&X)

	// 주어진 string X값을 슬라이스에 담는다
	word := strings.Split(X, "")

	// 슬라이스 길이 출력
	fmt.Println(len(word))
}