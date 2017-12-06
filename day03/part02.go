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

	maxX := input / 2
	maxY := input / 2

	a := make([][]int, maxY)

	for key, _ := range a {
		a[key] = make([]int, maxX)
	}

	pos   := Position{x: (input / 2) / 2, y: (input / 2) / 2}
	//start := pos
	last  := Position{}

	dir := "right"

	for i := 0; i < input; i++ {
		last.x = pos.x
		last.y = pos.y

		if i == 0 {
			a[pos.y][pos.x] = (i + 1)
			pos.x++
		} else {
			a[pos.y][pos.x] = getAdjSum(a, pos, maxX, maxY)
			if a[pos.y][pos.x] > input {
				break
			}
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

	return a[pos.y][pos.x]
}

func getAdjSum(a [][]int, pos Position, maxX int, maxY int) int {
	sum := 0

	if pos.y -1 >= 0 {
		sum += a[pos.y - 1][pos.x]
		if pos.x -1 >= 0 {
			sum += a[pos.y - 1][pos.x - 1]
		}
		if pos.x + 1 < maxX {
			sum += a[pos.y - 1][pos.x + 1]
		}
	}

	if pos.y + 1 < maxY {
		sum += a[pos.y + 1][pos.x]
		if pos.x - 1 >= 0 {
			sum += a[pos.y + 1][pos.x - 1]
		}
		if pos.x + 1 < maxX {
			sum += a[pos.y + 1][pos.x + 1]
		}
	}

	if pos.x - 1 >= 0 {
		sum += a[pos.y][pos.x - 1]
	}
	if pos.x + 1 < maxX {
		sum += a[pos.y][pos.x + 1]
	}

	return sum
}

func printArr(a [][]int) {
	for key, _ := range a {
		fmt.Println(a[key])
	}
}
