package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
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
			} else {
				for val, _ := range words {
					if len(word) == len(val) {
						// We need to check if this is an anagram
						if isAnagram(word, val) == true {
							valid = false
							break
						}
					}
				}
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

func isAnagram(a string, b string) bool {
	arrA := make([]string, len(a))
	arrB := make([]string, len(a))

	for k, val := range a {
		arrA[k] = string(val)
	}
	for k, val := range b {
		arrB[k] = string(val)
	}
	sort.Strings(arrA)
	sort.Strings(arrB)

	for i := 0; i < len(arrA); i++ {
		if arrA[i] != arrB[i] {
			return false;
		}
	}

	return true
}
