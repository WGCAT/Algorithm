// 예를 들어 array가 [1, (5, 2, 6, 3), 7, 4], i = 2, j = 5, k = 3이라면

// array의 2번째부터 5번째까지 자르면 [5, 2, 6, 3]입니다.
// 1에서 나온 배열을 정렬하면 [2, 3, (5), 6]입니다.
// 2에서 나온 배열의 3번째 숫자는 5입니다. [[2, 5, 3], [4, 4, 1], [1, 7, 3]]

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	response := solution(array, commands)

	fmt.Println(response)
}

func solution(array []int, commands [][]int) []int {
	// 결과값 담을 슬라이스 선언
	var result []int

	for _, value := range commands {
		var sliceSort []int
		if value[1] == len(array) {
			for _, sliceValue := range array[value[0]-1:] {
				sliceSort = append(sliceSort, sliceValue)
			}
			sort.Ints(sliceSort)
			result = append(result, sliceSort[value[2]-1])

		} else {
			for _, sliceValue := range array[value[0]-1 : value[1]] {
				sliceSort = append(sliceSort, sliceValue)
			}
			sort.Ints(sliceSort)
			result = append(result, sliceSort[value[2]-1])
		}
	}
	return result
}
