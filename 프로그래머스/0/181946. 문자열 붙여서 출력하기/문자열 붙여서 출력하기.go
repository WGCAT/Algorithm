package main

import (
	"fmt"
)

// 문자열 연산자로 붙이기
func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	str := s1 + s2
	fmt.Println(str)
}