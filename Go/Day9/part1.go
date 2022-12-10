package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	dirMap := map[string]image.Point{
		"C":  {0, 0},
		"R":  {1, 0},
		"L":  {-1, 0},
		"D":  {0, -1},
		"U":  {0, 1},
		"RD": {1, -1},
		"RU": {1, 1},
		"LD": {-1, -1},
		"LU": {-1, 1},
	}

	var headPosition image.Point = image.Point{0, 0}
	var tailPosition image.Point = image.Point{0, 0}
	var headPrevious image.Point = image.Point{0, 0}
	locations := make(map[image.Point]int)
	locations[image.Point{0, 0}] = 1

	for sc.Scan() {
		line := sc.Text()

		sections := strings.Fields(line)

		direction := sections[0]
		steps, _ := strconv.Atoi(sections[1])

		for i := 0; i < steps; i++ {
			headPrevious = headPosition
			headPosition = headPosition.Add(dirMap[direction])
			var withinRadius = false
			fmt.Println("Head", headPosition)
			for _, v := range dirMap {
				var tp = tailPosition
				if tp.Add(v) == headPosition {
					withinRadius = true
					break
				}
			}

			if !withinRadius {
				tailPosition = headPrevious
				locations[tailPosition] = 1
			}
			fmt.Println("Tail", tailPosition)
			fmt.Println()
		}

	}

	var sum = 0
	for _, v := range locations {
		sum += v
	}
	fmt.Println(sum)
}
