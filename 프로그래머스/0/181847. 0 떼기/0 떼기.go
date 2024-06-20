package algorithm

// 정수로 이루어진 문자열을 쪼개 슬라이스를 만들어 for문으로 왼쪽부터 0조건문 후 다시 문자열로 반환
import (
	"strings"
)

func solution(n_str string) string {
	// 값을 담을 문자열 슬라이스 선언
	var strPut []string
	// 주어진 문자열을 쪼개 담을 슬라이스 선언
	strSlice := strings.Split(n_str, "")
	// 인덱스카운트 숫자 선언 및 초기화
	index := 0
	// for range문으로 문자열이 담긴 슬라이스를 numStr로 담아서 돌림
	for _, numStr := range strSlice {
		// numStr에 0이 속해있다면
		if strings.Contains(numStr, "0") {
			// index 카운트 증가
			index++

			// numStr에 0이 속해있지 않다면
		} else {
			// strSlice의 인덱스값을 index로 넣어서 index의 뒷값들을 strPut에 담기
			strPut = strSlice[index:]
			// for문 종료
			break
		}

	}
	// 슬라이스를 문자열로 변환
	result := strings.Join(strPut, "")

	return result
}
