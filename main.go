// A ~ Z : 65 ~ 90, a ~ z : 97 ~ 122

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)

	var result string

	slice := strings.Split(s1, "")

	for _, v := range slice {
		if v >= "A" && v <= "Z" {
			word := string(v)
			word = strings.ToUpper(word)
			result = result + word
		} else {
			word := string(v)
			word = strings.ToLower(word)
			result = result + word
		}
	}
	fmt.Println(result)
}
