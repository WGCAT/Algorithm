package main

import "strings"

func solution(babbling []string) int {
	say := [4]string{"aya", "ye", "woo", "ma"}
	var result int
	var choice []string
	var ru string
	for _, v := range babbling {
		var flag bool = false
		for _, vv := range say {
			if strings.Contains(v, strings.Repeat(vv, 2)) {
				flag = true
				break
			}
		}
		if !flag {
			choice = append(choice, v)
		}

	}
	for _, v := range choice {
		ru = v
		for _, vv := range say {
			ru = strings.ReplaceAll(ru, vv, "|")
		}
		ru = strings.ReplaceAll(ru, "|", "")
		if ru == "" {
			result++
		}
	}
	return result
}
