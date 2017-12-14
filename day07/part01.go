package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	file := "input"
	fh, _ := os.Open(file)
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	parents   := make(map[string]int)
	children  := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) > 2 {
			// has children
			parents[fields[0]] = 1
			c := fields[3:]

			for _, v := range c {
				children[strings.Replace(v, ",", "", 1)] = 1
			}
		} else {
			parents[fields[0]] = 1
		}
	}

	var found string

	for p := range parents {
		_, ok := children[p]
		if !ok {
			found = p
			break
		}
	}

	fmt.Println(found)
}
