package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	file := "input"
	content, _ := ioutil.ReadFile(file)
	sum := 0
	step := len(content) / 2
	var current, next string
	
	for i, val := range content {
		current = string(val)
		
		if (i + step) >= len(content) {
			next = string(content[step - (len(content) - i)])
		} else {
			next = string(content[i + step])
		}

		if current == next {
			add, _ := strconv.Atoi(current)
			sum += add
		}
	}

	fmt.Println(sum)
}
