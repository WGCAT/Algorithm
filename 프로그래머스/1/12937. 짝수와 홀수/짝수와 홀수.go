package algorithm

// num이 짝수(0포함)이면 "Even", 홀수면 "Odd" 반환
func solution(num int) string {
	// 반환값을 담을 문자열 함수 선언
	var response string
	// num을 2로 나눈 나머지가 0일 경우
	if num%2 == 0 {
		// "Even" 반환
		response = "Even"
		// 나머지의 경우
	} else {
		// "Odd" 반환
		response = "Odd"
	}
	return response
}
