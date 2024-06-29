func solution(n int) int {
    result := 1
    for{
        if (6*result) % n != 0{
            result++
        }else {
            break
        }
    } 
        
    return result
}