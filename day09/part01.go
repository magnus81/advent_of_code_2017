package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file := "input"
	content, _ := ioutil.ReadFile(file)

	var current string
	skip   := false
	sum    := 0
	score  := 0
	inGarbage := false

	for _, val := range content {
		current = string(val)

		if !skip {
			switch current {
			case "{":
				if !inGarbage {
					score++
					sum += score
				}
			case "}":
				if !inGarbage {
					score--
				}
			case "<":
				inGarbage = true
			case ">":
				inGarbage = false
			case "!":
				skip = true
			}
		} else {
			skip = false
		}
	}

	fmt.Println(sum)
}
