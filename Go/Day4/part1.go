package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	contains := 0

	for sc.Scan() {
		line := sc.Text()
		sections := strings.Split(line, ",")

		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")

		start1, _ := strconv.Atoi(section1[0])
		finish1, _ := strconv.Atoi(section1[1])
		start2, _ := strconv.Atoi(section2[0])
		finish2, _ := strconv.Atoi(section2[1])

		if (start1 >= start2 && finish1 <= finish2) || (start2 >= start1 && finish2 <= finish1) {
			contains++
		}
	}
	fmt.Println(contains)
}
