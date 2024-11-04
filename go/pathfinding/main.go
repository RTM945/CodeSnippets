package main

import (
	"container/list"
	"fmt"
)

type coord struct {
	x, y int
}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func valid(c *coord, w, h int) bool {
	return 0 <= c.x && c.x < w && 0 <= c.y && c.y < h
}

func neighbors(m [][]int, c coord) []coord {
	h := len(m)
	w := len(m[0])
	res := make([]coord, 0)
	for _, d := range dirs {
		nc := &coord{c.x + d[0], c.y + d[1]}
		if valid(nc, w, h) {
			res = append(res, *nc)
		}
	}
	return res
}

func main() {
	m := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	start := coord{
		x: 2,
		y: 3,
	}
	goal := coord{
		x: 8,
		y: 0,
	}
	frontier := list.New()
	frontier.PushBack(start)
	came_from := map[coord]coord{}
	came_from[start] = coord{-1, -1}
	for frontier.Len() > 0 {
		current := frontier.Remove(frontier.Front()).(coord)
		if current.x == 9 && current.y == 0 {
			fmt.Print()
		}
		if current == goal {
			break
		}
		for _, next := range neighbors(m, current) {
			if _, ok := came_from[next]; !ok {
				frontier.PushBack(next)
				came_from[next] = current
			}
		}
	}
	fmt.Println(len(came_from))

	current := goal
	path := make([]coord, 0)
	for current != start {
		path = append(path, current)
		current = came_from[current]
	}
	path = append(path, start)
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i], "->")
	}
}
