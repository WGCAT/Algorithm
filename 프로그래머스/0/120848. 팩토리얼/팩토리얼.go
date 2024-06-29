func solution(n int) int {
    mul := 1
    result := 0
    for mul <= n {
        mul = mul * (result+1)
        result++
    }
    return result-1
}