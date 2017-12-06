package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Result struct {
	valid int
	invalid int
}

func main() {

	file := "input"
	fh, _ := os.Open(file)
	defer fh.Close()

	result := Result{0, 0}

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		words  := make(map[string]bool)
		valid  := true

		for _, word := range fields {
			_, set := words[word]
			if set {
				valid = false
				break
			}
			words[word] = true
		}

		if valid {
			result.valid++
		} else {
			result.invalid++
		}
	}

	fmt.Println(result.valid)
}
