package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	file := "input"
	fh, _ := os.Open(file)
	defer fh.Close()

	var largest, smallest string
	sum := 0

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		for k, val := range numbers {
			if k == 0 {
				largest = val
				smallest = val
			} else {
				if stoi(val) < stoi(smallest) {
					smallest = val
				}
				if stoi(val) > stoi(largest) {
					largest = val
				}
			}
		}
		sum = sum + (stoi(largest) - stoi(smallest))
	}

	fmt.Println(sum)
}

func stoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}
