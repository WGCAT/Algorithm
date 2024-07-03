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
		}
	}
	return result
}