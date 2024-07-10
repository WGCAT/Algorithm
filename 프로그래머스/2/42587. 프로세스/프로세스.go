func solution(priorities []int, location int) int {
	var order1 []int
	var order2 []int
	result := 1

	order1 = append(order1, priorities...)
	for i := 0; i < len(priorities); i++ {
		order2 = append(order2, i)
	}

	max := 0
	for _, v := range order1 {
		if v > max {
			max = v
		}
	}
	for len(order1) != 0 {
		a := order1[0]
		b := order2[0]
		order1 = append(order1[:0], order1[0+1:]...)
		order2 = append(order2[:0], order2[0+1:]...)
		if a == max && b == location {
			return result
		} else if a == max {
			result++
			max = 0
			for _, v := range order1 {
				if v > max {
					max = v
				}
			}
		} else {
			order1 = append(order1, a)
			order2 = append(order2, b)
		}
	}
	return result

}
