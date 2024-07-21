func solution(numbers []int64) []int {
	var binarySlice []string
	var binaryNum string

	for _, v := range numbers {
		// binaryNum = strconv.FormatInt(v, 2)
		binaryNum = binaryCal(v)
		num := 1
		for i := 1; i < 10; i++ {
			num = num * 2
			lenNode := num - 1
			if len(binaryNum) < lenNode {
				for len(binaryNum) != lenNode {
					binaryNum = "0" + binaryNum
				}
				break
			} else if len(binaryNum) == lenNode {
				break
			}
		}
		binarySlice = append(binarySlice, binaryNum)
	}

	i := 0
	result := make([]int, len(binarySlice))
	for _, v := range binarySlice {
		value := tree(v)
		result[i] = value
		i++
	}

	return result
}

func tree(number string) int {
	result := 0
	// if !strings.Contains(number, "0") || !strings.Contains(number, "1") {
	// 	result = 1
	// 	return result
	// }
	if containBinary(number) {
		result = 1
		return result
	}
	mid := len(number) / 2
	if number[mid] == '0' {
		return result
	}
	leftSub := tree(number[:mid])
	rightSub := tree(number[mid+1:])
	if leftSub == 1 && rightSub == 1 {
		result = 1
	}
	return result
}

func binaryCal(number int64) string {
	binary := ""
	for number != 0 {
		if number%2 == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		number = number / 2
	}
	return binary
}

func containBinary(number string) bool {
	result := false
	one := 0
	zero := 0
	for i := 0; i < len(number); i++ {
		if number[i] == '1' {
			one++
		} else {
			zero++
		}
	}
	if zero == len(number) || one == len(number) {
		result = true
	}
	return result
}
