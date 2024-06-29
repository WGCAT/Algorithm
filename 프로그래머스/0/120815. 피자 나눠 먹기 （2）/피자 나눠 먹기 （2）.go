func solution(n int) int {
    r := 1
    for (6*r) % n != 0 {
        r++
    }
    return r
}