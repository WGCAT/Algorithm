package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 제로
	// 정수 K가 주어진다. (1 ≤ K ≤ 100,000)
	// 이후 K개의 줄에 정수가 1개씩 주어진다
	var K, V int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &K)

	// 정수들을 담을 슬라이스 선언
	var stack []int

	// 인덱스 카운트 정수 선언
	index := 0

	// 정수K의 크기만큼 for문을 돌림
	for i := 0; i < K; i++ {
		fmt.Fscanln(reader, &V)
		// V가 0이라면
		if V == 0 {
			// stack의 index-1의 이전 인덱스값을 담는다
			stack = stack[:index-1]
			// 인덱스값 감소
			index--
			// 인덱스값 증가를 막기위해 continue
			continue
			// V가 0이 아니라면
		} else {
			// stack에 V를 담는다
			stack = append(stack, V)
		}
		// 인덱스값 증가
		index++
	}

	// 합한 값을 더할 변수 선언
	sum := 0
	// for range문으로 v에 stack의 값을 담아서 돌림
	for _, v := range stack {
		// sum에 v를 더함
		sum = sum + v
	}

	fmt.Println(sum)

}
