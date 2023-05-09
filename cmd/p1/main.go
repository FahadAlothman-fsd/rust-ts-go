package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Point struct {
	x int
	y int
}

func parseLine(line string) Point {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("this should never happen")
	}
	if parts[0] == "forward" {
		return Point{amount, 0}
	} else if parts[0] == "up" {
		return Point{x: 0, y: -amount}
	}
	return Point{0, amount}
}

func main() {
	lines := strings.Split(getInput(), "\n")
	pos := Point{0, 0}
	for _, lines := range lines {
		point := parseLine(lines)
		pos.x += point.x
		pos.y += point.y

	}
	fmt.Printf("point: %d %d", pos, pos.x * pos.y)
}