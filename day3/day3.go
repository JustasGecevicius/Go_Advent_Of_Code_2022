package day3

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getCorrectValue(val byte) int {
	intValue := int(val) 
	if (intValue >= 97) {
		return intValue - 96
	}
	return intValue - 38
}

func solution1() int {
	file, err := os.Open("day3/task.txt")
	if (err != nil) { panic(err) }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	dataSlice := make([]int, 53)
	resultsSlice := []int{}

	for scanner.Scan() {
		line := scanner.Bytes()
		slice1 := line[:int(math.Floor(float64(len(line))/2))]
		slice2 := line[int(math.Floor(float64(len(line))/2)):]

		for _, data := range slice1 {
			value := getCorrectValue(data)
			dataSlice[value] = 1
		}

		for _, data := range slice2 {
			value := getCorrectValue(data)
			if dataSlice[value] == 1 {
				dataSlice[value] = 2
			}
		}
		
		for i, data := range dataSlice {
			if data == 2 {
				resultsSlice = append(resultsSlice, i)	
			}
		}

		for i := range dataSlice {
			dataSlice[i] = 0
		}
	}
	total := 0
	for _, data := range resultsSlice	{
		total += data
	}
	return total
}

func solution2() int {
	file, err := os.Open("day3/task.txt")
	if (err != nil) { panic(err) }
	defer file.Close()
	scanner := bufio.NewScanner(file);
	groupLineCount := 0
	encounteredValues := make([]int, 53)
	results := []int{}
	for scanner.Scan() {
		line := scanner.Bytes()
		for _, data := range line {
			value := getCorrectValue(data)
			switch true {
			case groupLineCount == 0:
				encounteredValues[value] = 1	
			case groupLineCount == 1 && encounteredValues[value] == 1:
				encounteredValues[value] = 2	
			case groupLineCount == 2 && encounteredValues[value] == 2:
				encounteredValues[value] = 3	
			}
		}
		for i := range encounteredValues {
			if (encounteredValues[i] == 3) {
				results = append(results, i)
			}
		}
		if groupLineCount == 2 {
			for i := range encounteredValues {
				encounteredValues[i] = 0
			}
			groupLineCount = 0
		} else {
			groupLineCount++
		}
	}
	total := 0
	for _, data := range results {
		total += data		
	}
	return total
}

func Solution() {
	fmt.Printf("Day 3: Solution 1: %v, Solution 2: %v", solution1(),solution2())
	fmt.Println()
}