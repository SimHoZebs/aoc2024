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
	part2()
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

			if (absDiff > 3 || absDiff < 1) || (i > 1 && prevDiffSign != diffSign) {
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

func part2() {
	file, _ := os.Open("input")
	defer file.Close()

	var ans int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// fmt.Println(line)

		// if safe, i, j := isSafe(line, -1); safe {
		// 	ans++
		// } else if safe, _, _ := isSafe(line, i); safe {
		// 	ans++
		// } else if safe, _, _ := isSafe(line, j); safe {
		// 	ans++
		// } else {
		// 	fmt.Println("My attempt deems it Not Safe")
		// }

		for i := -1; i < len(line); i++ {
			if safe, _, _ := isSafe(line, i); safe {
				ans++
				break
			}
		}
	}

	fmt.Printf("part2: %d\n", ans)
}

func isSafe(arr []string, skip int) (bool, int, int) {
	var prevDiff = math.Inf(1)
	isSafe := true
	var i, j = 0, 1
	for j < len(arr) {

		if j == skip {
			j++
			continue
		} else if i == skip {
			i++
			j++
			continue
		}

		level1, _ := strconv.Atoi(arr[i])
		level2, _ := strconv.Atoi(arr[j])

		diff := float64(level1 - level2)
		diffSign := math.Signbit(diff)
		prevDiffSign := math.Signbit(prevDiff)
		absDiff := math.Abs(diff)

		if (absDiff > 3 || absDiff < 1) || (prevDiff != math.Inf(1) && prevDiffSign != diffSign) {
			isSafe = false
			break
		}

		i = j
		j++
		prevDiff = diff
	}

	return isSafe, i, j
}
