// for range문으로 앞에서부터 하나씩 더하다가 n보다 커지는순간 break

func solution(numbers []int, n int) int {
    // 더한 값들을 담을 변수 선언과 동시에 초기화
    total := 0
    // for range문으로 numbers의 값들을 sum에 담아서 돌림
    for _, sum := range numbers {
        // 만약 total값이 n보다 작거나 같다면
        if total <= n {
            // total에 sum을 더한다
            total = total + sum
        // 만약 total값이 n보다 크다면    
        }else {
            // break로 for문을 빠져나온다
            break 
        }
    }
    
    return total
}