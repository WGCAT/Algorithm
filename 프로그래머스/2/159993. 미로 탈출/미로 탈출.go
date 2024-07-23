func solution(maps []string) int {
	depth := len(maps)
	width := len(maps[0])
	var startX, startY int
	var leverX, leverY int
	var endX, endY int

	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
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

	leverTime := bfs(startX, startY, leverX, leverY, maps, depth, width)
	if leverTime == -1 {
		return -1
	}

	endTime := bfs(leverX, leverY, endX, endY, maps, depth, width)
	if endTime == -1 {
		return -1
	}

	return leverTime + endTime
}

func bfs(startX int, startY int, endX int, endY int, maps []string, depth int, width int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, depth)
	for i := 0; i < depth; i++ {
		visited[i] = make([]bool, width)
	}
	queue := [][]int{}
	start := []int{startX, startY, 0}
	queue = append(queue, start)

	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			visited[i] = append(visited[i], false)
		}
	}

	for len(queue) != 0 {
		x := queue[0][0]
		y := queue[0][1]
		time := queue[0][2]
		queue = append(queue[:0], queue[0+1:]...)

		if x == endX && y == endY {
			return time
		}
		if visited[x][y] {
			continue
		}
		visited[x][y] = true

		for _, v := range directions {
			newX := x + v[0]
			newY := y + v[1]

			if newX >= 0 && newX < depth && newY >= 0 && newY < width && !visited[newX][newY] && maps[newX][newY] != 'X' {
				add := []int{newX, newY, time + 1}
				queue = append(queue, add)
			}
		}
	}
	return -1

}
