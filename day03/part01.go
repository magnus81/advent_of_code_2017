package main

import (
	"fmt"
)

type Position struct {
	x int
	y int
}

func main() {

	input := 289326
	
	fmt.Println(getSteps(input))
}

func getSteps(input int) int {

	if input == 1 {
		return 0
	}

	a := make([][]int, input / 2)

	for key, _ := range a {
		a[key] = make([]int, input / 2)
	}

	pos   := Position{x: (input / 2) / 2, y: (input / 2) / 2}
	start := pos
	last  := Position{}

	dir := "right"

	for i := 0; i < input; i++ {
		last.x = pos.x
		last.y = pos.y

		a[pos.y][pos.x] = (i + 1)
		if i == 0 {
			pos.x++
		} else {
			switch dir {
				case "right":
					if a[pos.y - 1][pos.x] == 0 {
						pos.y--
						dir = "up"
					} else {
						pos.x++
					}
				case "left":
					if a[pos.y + 1][pos.x] == 0 {
						pos.y++
						dir = "down"
					} else {
						pos.x--
					}
				case "up":
					if a[pos.y][pos.x - 1] == 0 {
						pos.x--
						dir = "left"
					} else {
						pos.y--
					}
				case "down":
					if a[pos.y][pos.x + 1] == 0 {
						pos.x++
						dir = "right"
					} else {
						pos.y++
					}
			}
		}
	}

	return (getAbs(last.x, start.x) + getAbs(last.y, start.y))
}

func getAbs(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
