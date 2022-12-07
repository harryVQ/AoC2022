package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	winMap := make(map[string]string)
	winMap["A"] = "Y"
	winMap["B"] = "Z"
	winMap["C"] = "X"

	drawMap := make(map[string]string)
	drawMap["A"] = "X"
	drawMap["B"] = "Y"
	drawMap["C"] = "Z"

	pointsMap := make(map[string]int)
	pointsMap["X"] = 1
	pointsMap["Y"] = 2
	pointsMap["Z"] = 3

	points := 0

	for sc.Scan() {
		line := sc.Text()
		opponent := string(line[0])
		us := string(line[2])

		win := us == winMap[opponent]
		draw := us == drawMap[opponent]

		if win {
			points += 6 + pointsMap[us]
		} else if draw {
			points += 3 + pointsMap[us]
		} else {
			points += 0 + pointsMap[us]
		}
	}
	fmt.Println(points)
}
