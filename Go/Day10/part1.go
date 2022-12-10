package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkForCycle(c int, r int) int {
	cycles := []int{20, 60, 100, 140, 180, 220}
	for _, v := range cycles {
		if c == v {
			return v * r
		}
	}
	return 0
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	register := 1
	cycle := 1
	sum := 0

	for sc.Scan() {
		line := sc.Text()
		sections := strings.Fields(line)
		if sections[0] == "noop" {
			cycle++
			sum += checkForCycle(cycle, register)
		} else if sections[0] == "addx" {
			cycle++
			sum += checkForCycle(cycle, register)
			cycle++
			add, _ := strconv.Atoi(sections[1])
			register += add
			sum += checkForCycle(cycle, register)
		} else {
			fmt.Println("never HIT")
		}
	}
	fmt.Println(sum)
}
