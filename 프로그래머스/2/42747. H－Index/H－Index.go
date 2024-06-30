import (
	"sort"
)

func solution(citations []int) int {
	sort.Ints(citations)
	var result int
	var slice []int
	for h := 0; h < 1000; h++ {

		for i := 0; i < len(citations); i++ {
			slice = []int{}
			if citations[i]-h >= 0 {
				slice = append(slice, citations[i:]...)
				break
			}
		}
		if len(slice) < h {
			result = h - 1
			break
		}
	}
	return result
}