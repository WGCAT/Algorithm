package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 문자열 연산자로 붙이기
func main1() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	str := s1 + s2
	fmt.Println(str)
}

// 문자열을 배열에 담아 값들을 연결하는 strings.Join으로 출력
func main2() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	var word []string
	word = append(word, s1, s2)

	fmt.Println(strings.Join(word, ""))
}

// 문자열을 bytes.Buffer내에 있는 byte slice(buf)에 WriteString을 사용하여 인자로 받은 string 을 이어붙이기
func main3() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)

	fmt.Println(buffer.String())
}
