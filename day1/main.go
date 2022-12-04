package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	index    int
	calories int
}

func parser() []elf {
	var elves []elf
	var indexer int = -1
	f, err := os.Open("input_day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			var e elf
			e.index = indexer
			elves = append(elves, e)
            indexer++
		} else {
			calories, _ := strconv.Atoi(scanner.Text())
			elves[indexer].calories += calories
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    sort.Slice(elves, func(i, j int) bool {
        return elves[i].calories > elves[j].calories
    })
	return elves
}

func topThreeTotalCalories(elves []elf) int {
	var total int
	for i := 0; i < 3; i++ {
		total += elves[i].calories
	}
	return total
} 

func main() {
	elves := parser()
	fmt.Println(elves[0].calories)
	fmt.Println("Top 3 Calories: ", topThreeTotalCalories(elves))
}
