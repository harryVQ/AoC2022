package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var currentInt = 0

func addKey(s []string) []string {
	s = append(s, strconv.Itoa(currentInt))
	currentInt++
	return s
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var keys []string
	var sizeMap = make(map[string]int)

	for _, v := range lines {
		sections := strings.Fields(v)
		if sections[0] == "$" {
			if sections[1] == "cd" {
				if sections[2] == ".." {
					keys = keys[:len(keys)-1]
				} else if sections[2] == "/" {
					keys = []string{"/"}
				} else {
					keys = addKey(keys)
				}
			}
		} else if sections[0] == "dir" {
			continue
		} else {
			value, _ := strconv.Atoi(sections[0])
			for _, v := range keys {
				sizeMap[v] += value
			}
		}
	}

	var size = 70000000
	for _, value := range sizeMap {
		if value+70000000-sizeMap["/"] >= 30000000 && value < size {
			size = value
		}
	}

	fmt.Println(size)
}
