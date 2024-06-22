package main

import (
	"fmt"
	"strings"
)

func main() {
	var word []string
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	word = append(word, s1)
	word = append(word, s2)

	fmt.Println(strings.Join(word, ""))
}
