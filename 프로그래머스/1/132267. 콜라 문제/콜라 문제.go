func solution(a int, b int, n int) int {
	var result int
	for n >= a {
		c := (n / a) * b
		n = (n/a)*b + n%a
		result = result + c
	}
	return result
}
