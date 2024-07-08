func solution(n int, lost []int, reserve []int) int {
    s := make([]int, n)
    
    for i := 0; i < n; i++ {
        s[i] = 1 
    }
    for _, v := range reserve {
        s[v-1]++
    }
    for _, v := range lost {
        s[v-1]--
    }
	for i :=0; i< n; i++ {
        if s[i] == 0 {
            if i > 0 && s[i-1] == 2 {
                s[i-1]--
                s[i]++
            }else if i < n-1 && s[i+1] == 2 {
                s[i+1]--
                s[i]++
            }
        }
	}
    answer := 0
    for i :=0; i<n; i++ {
        if s[i] > 0 {
            answer++
        }
    }
	return answer
}