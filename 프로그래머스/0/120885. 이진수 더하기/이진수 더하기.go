import(
    "strconv"
)
func solution(bin1 string, bin2 string) string {
	var slice []string
	var sum int
	slice = append(slice, bin1, bin2)
	for _, v := range slice {
		for i := 0; i < len(v); i++ {
			if v[i] == 49 {
				sum = sum + 1<<(len(v)-1-i)
			}
		}
	}
	return strconv.FormatInt(int64(sum), 2)
}
