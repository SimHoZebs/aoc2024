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
		fmt.Printf("line: %v\n", line)

		var prevLevel int
		var prevDiffSign bool
		isSafe := true
		for i, str := range line {
			level, _ := strconv.Atoi(str)

			if i == 0 {
				prevLevel = level
				continue
			}

			diff := prevLevel - level
			diffSign := math.Signbit(float64(diff))
			absDiff := math.Abs(float64(diff))

			fmt.Println(prevLevel, "-", level, "=", diff)
			if absDiff > 3 || absDiff < 1 {
				isSafe = false
				break
			}

			if i == 1 {
				prevLevel = level
				prevDiffSign = diffSign
				continue
			}

			if prevDiffSign != diffSign {
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
