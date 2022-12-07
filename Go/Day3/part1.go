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

	sum := 0

	prioMap := make(map[string]int) // asciii map
	for i := 97; i < 123; i++ {
		prio := i - 96
		prioMap[string(i)] = prio
	}

	for i := 65; i < 91; i++ {
		prio := i - 64 + 26
		prioMap[string(i)] = prio
	}

	for sc.Scan() {
		line := sc.Text()
		length := len(line)
		compartmentSize := length / 2
		firstCompartment := line[0:compartmentSize]
		secondCompartMent := line[compartmentSize:]
		for i := 0; i < len(firstCompartment); i++ {
			if strings.Contains(secondCompartMent, string(firstCompartment[i])) {
				sum += prioMap[string(firstCompartment[i])]
				break
			}
		}
	}

	fmt.Println(sum)
}
