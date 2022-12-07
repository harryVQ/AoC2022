package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sum = 0

	prioMap := make(map[string]int) // ascii iimap
	for i := 97; i < 123; i++ {
		prio := i - 96
		prioMap[string(i)] = prio
	}

	for i := 65; i < 91; i++ {
		prio := i - 64 + 26
		prioMap[string(i)] = prio
	}

	var currentLine string
	var previousLine string
	var s []string = nil
	var step = 1

	for sc.Scan() {
		line := sc.Text()
		currentLine = line

		switch step {
		case 1:
			previousLine = currentLine
			s = nil
			step = 2
			continue
		case 2:
			for i := 0; i < len(line); i++ {
				if strings.Contains(previousLine, string(currentLine[i])) {
					s = append(s, string(currentLine[i]))
				}
			}
			previousLine = currentLine
			step = 3
			continue
		case 3:
			for i := 0; i < len(s); i++ {
				if strings.Contains(currentLine, s[i]) {
					sum += prioMap[s[i]]
					step = 1
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
