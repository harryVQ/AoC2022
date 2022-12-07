package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseCommand(commandString string) [3]rune {
	var command [3]rune
	fmt.Sscanf(commandString, "move %d from %c to %c", &command[0], &command[1], &command[2])
	return command
}

func main() {
	var commands [][3]rune

	data, _ := os.ReadFile("input.txt")
	crates := strings.Split(strings.Split(string(data), "\n\n")[0], "\n") // Crates top split
	keys := crates[len(crates)-1]

	stacks := map[rune]string{}
	for _, stack := range crates {
		for i, c := range stack {
			if unicode.IsLetter(c) {
				stacks[rune(keys[i])] += string(c)
			}
		}
	}

	for _, c := range strings.Split(strings.TrimSpace(strings.Split(string(data), "\n\n")[1]), "\n") { // command bottom split
		commands = append(commands, parseCommand(c))
	}

	for _, s := range commands {
		var qty int = int(s[0])
		var from, to rune = s[1], s[2]

		stacks[to] = stacks[from][:qty] + stacks[to]
		stacks[from] = stacks[from][qty:]
	}

	//squish keys
	keys = strings.ReplaceAll(keys, " ", "")

	topCrates := ""
	for _, key := range keys {
		topCrates += stacks[key][:1]
	}
	fmt.Println(topCrates)
}
