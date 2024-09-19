package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func solution1() string {
	file, err := os.Open("day5/task.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ySlice := [][]byte{}
	isSetupFinished := false
	for scanner.Scan() {
		res := scanner.Bytes()
		if (len(ySlice) == 0)	{
			for i := 0; i < (len(res) + 1) / 4; i++ {
				ySlice = append(ySlice, []byte{})
			}
		}


		if !isSetupFinished && string(res[1]) != "1"{
			// Do the setup
			for y := 0; y < len(res); y += 4 {
				if (res[y + 1] != 32) {
					ySlice[y / 4] = append(ySlice[y / 4], res[y + 1])
				}
			}
		} else if len(res) == 0 {
			// Reversing the slices
			for i := range ySlice {
				slices.Reverse(ySlice[i])
			}
		} else if isSetupFinished {
			// Do the manipulation
			dataSlice := [][]byte{{},{},{}}
			index := 0
			allowAppend := true
			for i := 5; i < len(res); i++ {
				if allowAppend && res[i] != 32 {
					dataSlice[index] = append(dataSlice[index], res[i])	
				}
				if (res[i] == 32) {
					if allowAppend {
						index += 1
					}
					allowAppend = !allowAppend
				}
			}
			count, _ := strconv.Atoi(string(dataSlice[0]))
			fromIndex, _ := strconv.Atoi(string(dataSlice[1]))
			toIndex, _ := strconv.Atoi(string(dataSlice[2]))
			for i := 0; i < count; i++ {
				removedElement := ySlice[fromIndex - 1][len(ySlice[fromIndex - 1]) - 1]
				ySlice[fromIndex - 1] = ySlice[fromIndex - 1][:len(ySlice[fromIndex - 1]) - 1]
				ySlice[toIndex - 1] = append(ySlice[toIndex - 1], removedElement)
			}
		}

			if !isSetupFinished && string(res[1]) == "1"{
				isSetupFinished = true
			}
	}
	solution1 := []byte{}
	for _, row := range ySlice {
		solution1 = append(solution1, row[len(row) - 1])
	}
	return string(solution1)
}

func solution2() string {
	file, err := os.Open("day5/task.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ySlice := [][]byte{}
	isSetupFinished := false
	for scanner.Scan() {
		res := scanner.Bytes()
		if (len(ySlice) == 0)	{
			for i := 0; i < (len(res) + 1) / 4; i++ {
				ySlice = append(ySlice, []byte{})
			}
		}

		if !isSetupFinished && string(res[1]) != "1"{
			// Do the setup
			for y := 0; y < len(res); y += 4 {
				if (res[y + 1] != 32) {
					ySlice[y / 4] = append(ySlice[y / 4], res[y + 1])
				}
			}
		} else if len(res) == 0 {
			// Reversing the slices
			for i := range ySlice {
				slices.Reverse(ySlice[i])
			}
		} else if isSetupFinished {
			// Do the manipulation
			dataSlice := [][]byte{{},{},{}}
			index := 0
			allowAppend := true
			for i := 5; i < len(res); i++ {
				if allowAppend && res[i] != 32 {
					dataSlice[index] = append(dataSlice[index], res[i])	
				}
				if (res[i] == 32) {
					if allowAppend {
						index += 1
					}
					allowAppend = !allowAppend
				}
			}
			count, _ := strconv.Atoi(string(dataSlice[0]))
			fromIndex, _ := strconv.Atoi(string(dataSlice[1]))
			toIndex, _ := strconv.Atoi(string(dataSlice[2]))
			removedSlice := ySlice[fromIndex - 1][len(ySlice[fromIndex - 1]) - count:]
			ySlice[fromIndex - 1] = ySlice[fromIndex - 1][:len(ySlice[fromIndex - 1]) - count]
			ySlice[toIndex - 1] = append(ySlice[toIndex - 1], removedSlice...)
		}

			if !isSetupFinished && string(res[1]) == "1"{
				isSetupFinished = true
			}
	}
	solution1 := []byte{}
	for _, row := range ySlice {
		solution1 = append(solution1, row[len(row) - 1])
	}
	return string(solution1)
}

func Solution() {
  fmt.Printf("Day 5: Solution 1: %v, Solution 2: %v", solution1(), solution2())
}