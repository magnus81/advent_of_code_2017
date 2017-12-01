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
	var current, next string
	
	for i, val := range content {
		current = string(val)
		
		if (i + 1) == len(content) {
			next = string(content[0])
		} else {
			next = string(content[i + 1])
		}

		if current == next {
			add, _ := strconv.Atoi(current)
			sum += add
		}
	}

	fmt.Println(sum)
}
