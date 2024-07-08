package main

import (
	"fmt"
)

// [0] 0
// [3, 4] 2
// [1, 2, 3, 5, 6, 7, 10, 11] 5
// [3, 5, 11, 6, 1, 5, 3, 3, 1, 41] 5
// [1, 11, 111, 1111] 3

// h번 이상 인용된 논문이 h편 이상
// h의 최댓값이 반환

func main() {
	// lost := []int{1, 2, 3, 4, 5, 6}
	// reserve := []int{1, 2, 3}
	// r := solution(10, lost, reserve)
	answers := []int{1, 2, 3, 4, 5, 6}
	r := solution(answers)
	fmt.Println(r)

}

// {4, 6}
// {8, 9}
//
//	func solution(n int, lost []int, reserve []int) int {
//		for r := 0; r < len(reserve); r++ {
//			if len(reserve) == 0 {
//				break
//			}
//			for l := 0; l < len(lost); l++ {
//				if len(reserve) == 0 {
//					break
//				} else if reserve[r] == lost[l] {
//					lost = append(lost[:l], lost[l+1:]...)
//					reserve = append(reserve[:r], reserve[r+1:]...)
//					l--
//				} else if (reserve[r]-1) == lost[l] || (reserve[r]+1) == lost[l] {
//					lost = append(lost[:l], lost[l+1:]...)
//					reserve = append(reserve[:r], reserve[r+1:]...)
//					l--
//				}
//			}
//		}
//		return n - len(lost)
//	}
func solution(answers []int) []int {
	var student1 = [5]int{1, 2, 3, 4, 5}
	var student2 = [8]int{2, 1, 2, 3, 2, 4, 2, 5}
	var student3 = [10]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
	scores := make([]int, 3)
	var result []int
	for i := 0; i < len(answers); i++ {
		answer := answers[i]
		if answer == student1[i%len(student1)] {
			scores[0]++
		}
		if answer == student2[i%len(student2)] {
			scores[1]++
		}
		if answer == student3[i%len(student3)] {
			scores[2]++
		}
	}
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	for i := 0; i < len(scores); i++ {
		if scores[i] == max {
			result = append(result, i+1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}
