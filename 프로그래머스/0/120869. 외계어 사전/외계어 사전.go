import (
	"sort"
	"strings"
)
func solution(spell []string, dic []string) int {
	result := 2
	sort.Strings(spell)
	spellStr := strings.Join(spell, "")
	for _, v := range dic {
		slice := strings.Split(v, "")
		sort.Strings(slice)
		dicStr := strings.Join(slice, "")
		if strings.Contains(dicStr, spellStr) {
			result = 1
		}
	}
	return result
}