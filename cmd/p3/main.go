package main

import (
	"fmt"
	"strings"
)

func get_input() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	treeCount := 0
	for r, line := range strings.Split(get_input(), "\n") {
		if string(line[(r*3)%len(line)]) == "#" {
			treeCount += 1
		}
	}
	fmt.Printf("tree count: %v\n", treeCount)
}
