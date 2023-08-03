package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	data := make([]byte, 100)
	_, err = file.Read(data)
	check(err)

	text := string(data)

	groups := strings.Split(text, "\n\n")
	max := 0
	for _, group := range groups {
		calories := strings.Split(group, "\n")

		sum := 0
		for _, calorie := range calories {
			num, err := strconv.Atoi(calorie)
			if err != nil {
			} else {
				sum += num
			}
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Println("max", max)
}
