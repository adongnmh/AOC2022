package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func groupAndSumCalories(calorieInput string) []int {

	groupedCalories := strings.Split(calorieInput, "\n")

	elfNumber := 0
	var elfCaloriesMapping []int
	elfCaloriesMapping = append(elfCaloriesMapping, 0)

	for i, v := range groupedCalories {
		if v == "" && i != len(groupedCalories) {
			elfNumber++
			elfCaloriesMapping = append(elfCaloriesMapping, 0)
			continue
		}

		calories, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		elfCaloriesMapping[elfNumber] += calories
	}
	return elfCaloriesMapping
}

func puzzle1(inputData string) int {
	groupedCalories := groupAndSumCalories(inputData)
	sort.Ints(groupedCalories)
	topElf := groupedCalories[len(groupedCalories)-1]
	return topElf
}

func puzzle2(inputData string) int {
	groupedCalories := groupAndSumCalories(inputData)
	sort.Ints(groupedCalories)
	topThree := groupedCalories[len(groupedCalories)-1] + groupedCalories[len(groupedCalories)-2] + groupedCalories[len(groupedCalories)-3]
	return topThree
}

func main() {
	fmt.Println("AOC 2022 - Day 1")
	inputData, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		log.Fatal("No Food")
	}

	fmt.Println(puzzle1(string(inputData)))
	fmt.Println(puzzle2(string(inputData)))

}
