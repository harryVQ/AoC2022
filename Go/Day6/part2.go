package main

import (
	"fmt"
	"os"
)

func unique(s []string) bool {
	occurances := make(map[string]int)
	for i := 0; i < len(s); i++ {
		occurances[string(s[i])]++
	}

	for _, v := range occurances {
		if v > 1 {
			return false
		}
	}

	return true
}

func main() {
	data, _ := os.ReadFile("input.txt")

	var buffer []string

	for i := 0; i < len(data); i++ {
		if i < 13 {
			buffer = append(buffer, string(data[i]))
			continue
		}

		buffer = append(buffer, string(data[i]))

		if unique(buffer) {
			fmt.Println(i + 1)
			fmt.Println(buffer)
			break
		}

		buffer = buffer[1:]
	}
}
