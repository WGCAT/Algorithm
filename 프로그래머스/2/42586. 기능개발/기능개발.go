func solution(progresses []int, speeds []int) []int {
	var day []int
	var result []int
	for i := 0; i < len(progresses); i++ {
		if (100-progresses[i])%speeds[i] != 0 {
			day = append(day, (100-progresses[i])/speeds[i]+1)
		} else {
			day = append(day, (100-progresses[i])/speeds[i])
		}
	}
	i := 0
	c := 1
	for j := 0; j < len(day); j++ {
		if j == 0 {
			result = append(result, c)
		} else if day[j] > day[j-1] {
			i++
			c = 1
			result = append(result, c)
		} else {
			c++
			day[j] = day[j-1]
			result = append(result[:i], c)
		}

	}
	return result
}