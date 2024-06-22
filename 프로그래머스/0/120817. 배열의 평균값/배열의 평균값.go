// 정수 배열 numbers에 들어오는 값들을 더한 후 length로 나눈다

func solution(numbers []int) float64 {
    // 합계 int 변수 선언
    var sum int
    
    // for range문으로 numbers의 값을 정수v에 담아 돌림
    for _, v := range numbers {
        sum = sum + v
    }
    
    // 더해진 합계값과 numbers의 길이를 미리 float64로 형변환 시킨 후 평균값을 계산
    average := float64(sum)/float64(len(numbers))
    
    return average
}