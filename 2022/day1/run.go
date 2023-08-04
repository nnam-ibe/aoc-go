package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitOnEmptyLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// split on empty line
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}
	// if at end of file, return remaining data
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	// request more data
	return 0, nil, nil
}

func main() {
	file, err := os.Open("2022/day1/input.txt")
	check(err)

	defer file.Close()

	elves := []int{}
	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnEmptyLine)
	for scanner.Scan() {
		group := scanner.Text()
		calories := strings.Split(group, "\n")

		sum := 0

		for _, calorie := range calories {
			num, err := strconv.Atoi(calorie)
			check(err)

			sum += num
		}

		elves = append(elves, sum)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	fmt.Println("Part 1: ", elves[0])
	fmt.Println("Part 2: ", elves[0]+elves[1]+elves[2])
}
