func solution(n int) int {
    result := 1
    for{
        if (6*result) % n == 0{
            return result
        }
        result++
    } 
}