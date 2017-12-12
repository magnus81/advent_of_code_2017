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
	inGarbage := false

	for _, val := range content {
		current = string(val)

		if !skip {
			switch current {
			case "<":
				if inGarbage {
					sum++
				}
				inGarbage = true
			case ">":
				inGarbage = false
			case "!":
				skip = true
			default:
				if inGarbage {
					sum++
				}
			}
		} else {
			skip = false
		}
	}

	fmt.Println(sum)
}
