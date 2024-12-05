package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	var list1 []int64
	var list2 []int64

	scanner := bufio.NewScanner(file)

	//linear
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		item1, _ := strconv.ParseInt(line[0], 10, 64)
		list1 = append(list1, item1)
		item2, _ := strconv.ParseInt(line[1], 10, 64)
		list2 = append(list2, item2)
	}

	//linear
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	//linear
	var ans float64
	for i := range list1 {
		ans += math.Abs(float64(list1[i] - list2[i]))
	}
	fmt.Printf("part1: %.f\n", ans)

	// O(4n) = O(n)
}

func part2() {
	file, _ := os.Open("input")
	defer file.Close()

	var list1 []int64
	var list2 []int64

	scanner := bufio.NewScanner(file)

	//linear
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		item1, _ := strconv.ParseInt(line[0], 10, 64)
		list1 = append(list1, item1)
		item2, _ := strconv.ParseInt(line[1], 10, 64)
		list2 = append(list2, item2)
	}

	//linear
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	//linear
	freqMap := make(map[int64]int64, len(list2))
	for _, item := range list2 {
		freqMap[item]++
	}

	//linear
	var ans int64 = 0
	for _, item := range list1 {
		ans += int64(item) * freqMap[item]
	}

	fmt.Printf("part2: %d\n", ans)

}
