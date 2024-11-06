package main

import (
	"fmt"
)

type coord struct {
	x, y int
}

type coordPriority struct {
	c        coord
	priority int
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

func cost(m [][]int, from, to coord) int {
	return m[to.y][to.x]
}

func main() {
	m := [][]int{
		{2, 3, 1, 5, 1, 1, 2, 3, 0, 1, 1},
		{1, 2, 4, 3, 4, 3, 1, 2, 1, 3, 2},
		{3, 1, 5, 2, 1, 2, 2, 4, 1, 4, 3},
		{1, 2, 0, 1, 1, 1, 1, 1, 1, 2, 1},
	}

	start := coord{
		x: 2,
		y: 3,
	}
	goal := coord{
		x: 8,
		y: 0,
	}
	frontier := NewPriorityQueue(func(a, b coordPriority) bool {
		return b.priority > a.priority
	})
	frontier.Add(coordPriority{
		c:        start,
		priority: 0,
	})
	came_from := map[coord]coord{}
	came_from[start] = coord{-1, -1}
	cost_so_far := map[coord]int{}
	cost_so_far[start] = 0
	for frontier.Len() > 0 {
		current := frontier.Remove()
		if current.c == goal {
			break
		}
		for _, next := range neighbors(m, current.c) {
			new_cost := cost_so_far[current.c] + cost(m, current.c, next)
			_, ok := cost_so_far[next]

			if !ok || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				frontier.Add(coordPriority{
					c:        next,
					priority: new_cost,
				})
				came_from[next] = current.c
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
