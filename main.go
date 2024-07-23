package main

import (
	"fmt"
)

func main() {
	w := []string{"SOOOL", "XXXXO", "OOOOO", "OXXXX", "OOOOE"}
	r := solution(w)
	fmt.Println(r)
}
func solution(maps []string) int {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, -1, 1}
	var startX, startY int
	var leverLen, leverX, leverY int
	var endLen, endX, endY int

	for i := 0; i < len(maps); i++ {
		for j := 0; j < len(maps[0]); j++ {
			if maps[i][j] == 'S' {
				startX = i
				startY = j
			} else if maps[i][j] == 'L' {
				leverX = i
				leverY = j
			} else if maps[i][j] == 'E' {
				endX = i
				endY = j
			}
		}
	}


	arr := findPoint(mapSlice, 1)
	x := arr[1]
	y := arr[0]
	var check [][]bool{}

	return 0
}

func bfs(maps []string, startX int, startY int, endX int, endY int) {
	visited := make([][]bool, len(maps), len(maps[0]))
	for i:=0; i<len(visited); i++ {
		for j:=0; j<len(visited[i]); j++ {
			visited[i][j] = false
		}
	}
	var queue [][]int
	var startPoint []int{startX, startY, 0}
	queue = append(queue, startPoint)



}