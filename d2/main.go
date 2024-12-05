package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	file, _ := os.Open("input")
	defer file.Close()

	var ans int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		var prevDiffSign bool
		isSafe := true
		prevLevel, _ := strconv.Atoi(line[0])
		for i := 1; i < len(line); i++ {
			level, _ := strconv.Atoi(line[i])

			diff := float64(prevLevel - level)
			diffSign := math.Signbit(diff)
			absDiff := math.Abs(diff)

			if (absDiff > 3 || absDiff < 1) ||
				(i > 1 && prevDiffSign != diffSign) {
				isSafe = false
				break
			}

			prevLevel = level
			prevDiffSign = diffSign

		}

		if isSafe {
			ans++
		}

	}
	fmt.Printf("part1: %d\n", ans)
}
