package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func rps(inputData string) int {
	totalScore := 0
	inputSlice := strings.Split(inputData, "\n")

	scoreHashMap := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	for _, v := range inputSlice {
		totalScore += scoreHashMap[v]
	}
	return totalScore
}

func rps2(inputData string) int {
	totalScore := 0
	inputSlice := strings.Split(inputData, "\n")

	scoreHashMap := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	for _, v := range inputSlice {
		totalScore += scoreHashMap[v]
	}
	return totalScore
}

func main() {
	fmt.Println("AOC 2022 - Day 2")
	inputData, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rps(string(inputData)))
	fmt.Println(rps2(string(inputData)))
}
