package main

import (
	"fmt"
	"os"
)

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 文件读取
func readMaze(filename string) [][]int {
	// 文件加载
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 读取行列
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 构建二维数组
	maze := make([][]int, row)
	for i := range maze {
		// 换行会被读成0 所以这边扔掉每行行末的读取
		fmt.Fscanf(file, "%d")
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func walk(maze [][]int, start, end point) [][]int {
	// 初始化探索地图
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	// 初始化探索队列
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		// 发现终点 退出
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			// 如果越界或者撞墙
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// 如果越界或者已探索
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 如果回到起点
			if next == start {
				continue
			}
			// 探索点位的步数在当前步数上加一
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			// 探索点位加入队列
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	// 迷宫读取
	fmt.Println("迷宫地图：")
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	// 迷宫路径
	fmt.Println("探索地图：")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
