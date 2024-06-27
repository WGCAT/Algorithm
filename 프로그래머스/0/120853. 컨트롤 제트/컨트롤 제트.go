import (
	"strconv"
	"strings"
)
func solution(s string) int {
	result := 0
	slice := strings.Split(s, " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] != "Z" {
			numPlus, _ := strconv.ParseInt(slice[i], 10, 0)
			result = result + int(numPlus)
		} else {
			numMinus, _ := strconv.ParseInt(slice[i-1], 10, 0)
			result = result - int(numMinus)
		}
	}
	return result

}