package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 1, 3, 1, 4, 1, 2}
	r := solution(a)
	fmt.Println(r)

}

func solution(topping []int) int {
	cs := make([]int, len(topping))
	bro := make([]int, len(topping))
	csCount := 0
	broCount := 0
	result := 0
	for _, v := range topping {
		if cs[v] == 0 {
			csCount++
		}
		cs[v]++
	}
	for _, v := range topping {
		if bro[v] == 0 {
			broCount++
		}
		bro[v]++
		cs[v]--
		if cs[v] == 0 {
			csCount--
		}
		if csCount == broCount {
			result++
		}

	}

	return result
}
