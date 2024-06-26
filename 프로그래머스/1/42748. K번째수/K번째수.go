package main

import "sort"

func solution(array []int, commands [][]int) []int {
	// 결과값 담을 슬라이스 선언
	var result []int

	// for range문을 이용하여 commands의 값을 value에 담아 돌림
	for _, value := range commands {
		var sliceSort []int
		var slice []int
		// value[1]의 값이 array의 길이와 같을 경우 (같으면 slice의 범위를 벗어남)
		if value[1] == len(array) {
			// array를 잘라 slice에 담는다
			slice = array[value[0]-1:]
			// slice의 값을 sliceValue에 담아 돌림
			for _, sliceValue := range slice {
				// sliceSort에 sliceValue의 값 추가
				sliceSort = append(sliceSort, sliceValue)
			}
			// sliceSort 오름차순 정렬
			sort.Ints(sliceSort)
			// result에 sliceSort의 인덱스값 추가
			result = append(result, sliceSort[value[2]-1])
			// 같지 않다면
		} else {
			// array를 잘라 slice에 담는다
			slice = array[value[0]-1 : value[1]]
			// slice의 값을 sliceValue에 담아 돌림
			for _, sliceValue := range slice {
				// sliceSort에 sliceValue의 값 추가
				sliceSort = append(sliceSort, sliceValue)
			}
			// sliceSort 오름차순 정렬
			sort.Ints(sliceSort)
			// result에 sliceSort의 인덱스값 추가
			result = append(result, sliceSort[value[2]-1])
		}
	}
	return result
}
