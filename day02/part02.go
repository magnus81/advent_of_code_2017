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

	sum := 0
	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		for k, val := range numbers {
			
			for ki, vali := range numbers {
				if k == ki {
					continue
				} else if stoi(val) > stoi(vali) {
					if stoi(val) % stoi(vali) == 0 {
						sum = sum + (stoi(val) / stoi(vali))
						continue
					}
				}
			}
		}
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
