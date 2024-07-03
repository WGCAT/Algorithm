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
	if scores[0] > scores[1] {
		if scores[0] > scores[2] {
			result = append(result, 1)
		} else if scores[0] < scores[2] {
			result = append(result, 3)
		} else if scores[0] == scores[2] {
			result = append(result, 1, 3)
		}
	} else if scores[0] < scores[1] {
		if scores[1] > scores[2] {
			result = append(result, 2)
		} else if scores[1] < scores[2] {
			result = append(result, 3)
		} else if scores[1] == scores[2] {
			result = append(result, 2, 3)
		}
	} else if scores[0] == scores[1] {
		if scores[0] > scores[2] {
			result = append(result, 1, 2)
		} else if scores[0] < scores[2] {
			result = append(result, 2)
		} else if scores[0] == scores[2] {
			result = append(result, 1, 2, 3)
		}
	}
	return result
}