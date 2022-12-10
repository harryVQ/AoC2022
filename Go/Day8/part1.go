package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func parseInputData(data []byte) map[image.Point]int {
	m := make(map[image.Point]int)
	for y, str := range strings.Fields(string(data)) {
		for x, r := range str {
			m[image.Point{x, y}] = int(r - 48) // rune to int
		}
	}
	return m
}

func main() {
	data, _ := os.ReadFile("input.txt")

	var treeMap = parseInputData(data)
	directions := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var sum = 0
	for treePoint, treeHeight := range treeMap {
		var visible = 0

		for _, direction := range directions {
			// Move from point outwards in each direction, visable if map out of bounds
			for nextPoint := treePoint.Add(direction); ; nextPoint = nextPoint.Add(direction) {
				if _, valid := treeMap[nextPoint]; !valid { //Hit an edge, must be visable
					visible = 1
					break
				}

				if treeMap[nextPoint] >= treeHeight {
					// If value of outward point is larger then not visable from edge
					break
				}
			}
		}

		sum += visible
	}

	fmt.Println(sum)
}
