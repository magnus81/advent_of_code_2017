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

	registers := make(map[string]int)

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		isTrue := false

		_, set := registers[fields[0]]

		if !set {
			registers[fields[0]] = 0
		}

		switch fields[5] {
			case "==":
				if registers[fields[4]] == stoi(fields[6]) {
					isTrue = true
				}
			case "<":
				if registers[fields[4]] < stoi(fields[6]) {
					isTrue = true
				}
			case ">":
				if registers[fields[4]] > stoi(fields[6]) {
					isTrue = true
				}
			case ">=":
				if registers[fields[4]] >= stoi(fields[6]) {
					isTrue = true
				}
			case "<=":
				if registers[fields[4]] <= stoi(fields[6]) {
					isTrue = true
				}
			case "!=":
				if registers[fields[4]] != stoi(fields[6]) {
					isTrue = true
				}
		}

		if isTrue {
			switch fields[1] {
				case "inc":
					registers[fields[0]] += stoi(fields[2])
				case "dec":
					registers[fields[0]] -= stoi(fields[2])
			}
		}
	}

	fmt.Println(getLargest(registers))
}

func getLargest(m map[string]int) int {
	highest := 0
	for _, value := range m {
		if value > highest {
			highest = value
		}
	}

	return highest
}

func stoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}
