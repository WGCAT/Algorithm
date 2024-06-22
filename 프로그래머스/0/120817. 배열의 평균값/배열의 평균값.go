package algorithm

// 정수 배열 numbers에 들어오는 값들을 더한 후 length로 나눈다

func solution1(numbers []int) float64 {
	// 합계 int 변수 선언
	var sum int

	// for range문으로 numbers의 값을 정수v에 담아 돌림
	for _, v := range numbers {
		sum = sum + v
	}

	// 더해진 합계값과 numbers의 길이를 미리 float64로 형변환 시킨 후 평균값을 계산
	average := float64(sum) / float64(len(numbers))

	return average
}

// numbers의 길이가 짝수이면 배열의 인덱스최소값과 인덱스최대값을 더한 값을 2로 나눈값
// numbers의 길이가 홀수이면 배열의 중간값
func solution2(numbers []int) float64 {
	var result float64
	if len(numbers)%2 == 0 {
		result = float64((numbers[0] + numbers[len(numbers)-1])) / float64(2)
	} else {
		result = float64(numbers[numbers[len(numbers)-1]/2])
	}
	return result

}
