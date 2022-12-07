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

	typeMap := make(map[string]map[string]string)
	typeMap["A"] = make(map[string]string)
	typeMap["B"] = make(map[string]string)
	typeMap["C"] = make(map[string]string)
	typeMap["A"]["X"] = "C"
	typeMap["A"]["Y"] = "A"
	typeMap["A"]["Z"] = "B"
	typeMap["B"]["X"] = "A"
	typeMap["B"]["Y"] = "B"
	typeMap["B"]["Z"] = "C"
	typeMap["C"]["X"] = "B"
	typeMap["C"]["Y"] = "C"
	typeMap["C"]["Z"] = "A"

	pointsMap := make(map[string]int)
	pointsMap["A"] = 1
	pointsMap["B"] = 2
	pointsMap["C"] = 3

	outcomeMap := make(map[string]int)
	outcomeMap["X"] = 0
	outcomeMap["Y"] = 3
	outcomeMap["Z"] = 6

	points := 0

	for sc.Scan() {
		line := sc.Text()
		opponent := string(line[0])
		outcome := string(line[2])

		points += outcomeMap[outcome] + pointsMap[typeMap[opponent][outcome]]

	}
	fmt.Println(points)
}
