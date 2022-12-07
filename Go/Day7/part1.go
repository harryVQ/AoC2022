package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var currentInt = 0 // used as unique indexs for directory names

func addKey(s []string) []string { // We dont carea bout dir name so use increasing ints instead to avoid duplicates
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

	var sum = 0
	for _, value := range sizeMap {
		if value <= 100000 {
			sum += value
		}
	}

	fmt.Println(sum)
}
