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

	locations := make(map[image.Point]int)
	locations[image.Point{0, 0}] = 1
	rope := []image.Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	for sc.Scan() {
		line := sc.Text()

		sections := strings.Fields(line)

		direction := sections[0]
		steps, _ := strconv.Atoi(sections[1])

		for i := 0; i < steps; i++ {
			var previousPosition = image.Point{0, 0}
			fmt.Println(rope)
			for k := 0; i < len(rope); i++ {
				var withinRadius = false
				if k == 0 {
					rope[k] = rope[k].Add(dirMap[direction])
					continue
				}
				for _, v := range dirMap {
					if k != 0 {
						if rope[k].Add(v) == rope[k-1] {
							withinRadius = true
							break
						}
					}
				}
				if !withinRadius {
					var prev = rope[k]
					rope[k] = previousPosition
					previousPosition = prev
				} else {
					previousPosition = rope[k]
				}
			}
		}

	}

	var sum = 0
	for _, v := range locations {
		sum += v
	}
	fmt.Println(sum)
}
