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

	parents := []string{}
	childs  := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "->")
		
		if len(s) > 1 {
			// this one has childs
			tmp := strings.Split(s[1], ", ")
			for _, v := range tmp {
				childs = append(childs, strings.TrimSpace(v))
			}
			tmp_2 := strings.Split(s[0], " ")
			parents = append(parents, tmp_2[0])
		} else {
			tmp := strings.Split(s[0], " ")
			parents = append(parents, tmp[0])
		}
	}

	var found string

	for _, v := range parents {
		if !inArray(childs, v) {
			found = v
			break
		}
	}

	fmt.Println(found)
}

func inArray(a []string, val string) bool {
	for _, v := range a {
		if val == v {
			return true
		}
	}

	return false
}
