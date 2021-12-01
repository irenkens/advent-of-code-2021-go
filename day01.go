package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Day1() {
	content, _ := ioutil.ReadFile("./input/day01.txt")
	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	var numbers []int

	for _, a := range lines {
		number, _ := strconv.Atoi(a)
		numbers = append(numbers, number)
	}

	fmt.Println("---- Day 1 ----")
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}

func part1(numbers []int) (increased int) {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1] > numbers[i] {
			increased++
		}
	}
	return increased
}

func part2(numbers []int) (increased int) {
	for i := 0; i < len(numbers)-3; i++ {
		group1 := numbers[i] + numbers[i+1] + numbers[i+2]
		group2 := numbers[i+1] + numbers[i+2] + numbers[i+3]
		if group2 > group1 {
			increased++
		}
	}
	return increased
}
