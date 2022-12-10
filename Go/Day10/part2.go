package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkForCycle(c int, r int, screenMap map[int][]string) map[int][]string {
	var key = c / 40
	var relativeC = c % 40
	if relativeC == (r-1) || relativeC == r || relativeC == (r+1) {
		screenMap[key] = append(screenMap[key], "█")
	} else {
		screenMap[key] = append(screenMap[key], ".")
	}
	return screenMap
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	screen := make(map[int][]string)

	register := 1
	cycle := 0
	for sc.Scan() {
		line := sc.Text()
		sections := strings.Fields(line)
		if sections[0] == "noop" {
			screen = checkForCycle(cycle, register, screen)
			cycle++
		} else if sections[0] == "addx" {
			screen = checkForCycle(cycle, register, screen)
			cycle++
			screen = checkForCycle(cycle, register, screen)
			cycle++
			add, _ := strconv.Atoi(sections[1])
			register += add
		} else {
			fmt.Println("never HIT")
		}
	}
	fmt.Println(screen[0])
	fmt.Println(screen[1])
	fmt.Println(screen[2])
	fmt.Println(screen[3])
	fmt.Println(screen[4])
	fmt.Println(screen[5])
}
