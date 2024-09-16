package day4

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func solution1() int {
	file, err := os.Open("day4/task.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file);

	count := 0
	
	for scanner.Scan() {
		line := scanner.Bytes()
		commaIndex := slices.IndexFunc(line, func(x byte) bool { return x == 44 })
		if (commaIndex == -1) { panic("ERROR"); }

		elf1Slice, elf2Slice := line[0:commaIndex], line[commaIndex + 1:]			

		elf1DashIndex := slices.IndexFunc(elf1Slice, func(x byte) bool { return x == 45 })
		elf2DashIndex := slices.IndexFunc(elf2Slice, func(x byte) bool { return x == 45 })

		elf1RangeStart, _:= strconv.Atoi(string(elf1Slice[0:elf1DashIndex]))
		elf1RangeEnd, _:= strconv.Atoi(string(elf1Slice[elf1DashIndex + 1:]))
		elf2RangeStart, _:= strconv.Atoi(string(elf2Slice[0:elf2DashIndex]))
		elf2RangeEnd, _:= strconv.Atoi(string(elf2Slice[elf2DashIndex + 1:]))
	
		if elf1RangeStart <= elf2RangeStart && elf1RangeEnd >= elf2RangeEnd {
			count++
		} else if elf1RangeStart >= elf2RangeStart && elf1RangeEnd <= elf2RangeEnd {
			count++
		}
	}
	return count
}

func solution2() int {
	file, err := os.Open("day4/task.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file);

	count := 0
	
	for scanner.Scan() {
		line := scanner.Bytes()
		commaIndex := slices.IndexFunc(line, func(x byte) bool { return x == 44 })
		if (commaIndex == -1) { panic("ERROR"); }

		elf1Slice, elf2Slice := line[0:commaIndex], line[commaIndex + 1:]			

		elf1DashIndex := slices.IndexFunc(elf1Slice, func(x byte) bool { return x == 45 })
		elf2DashIndex := slices.IndexFunc(elf2Slice, func(x byte) bool { return x == 45 })

		elf1RangeStart, _:= strconv.Atoi(string(elf1Slice[0:elf1DashIndex]))
		elf1RangeEnd, _:= strconv.Atoi(string(elf1Slice[elf1DashIndex + 1:]))
		elf2RangeStart, _:= strconv.Atoi(string(elf2Slice[0:elf2DashIndex]))
		elf2RangeEnd, _:= strconv.Atoi(string(elf2Slice[elf2DashIndex + 1:]))
	
		if elf1RangeStart <= elf2RangeStart && elf1RangeEnd >= elf2RangeEnd {
			count++
		} else if elf1RangeStart >= elf2RangeStart && elf1RangeEnd <= elf2RangeEnd {
			count++
		} else if elf1RangeStart >= elf2RangeStart && elf1RangeStart <= elf2RangeEnd {
			count++
		} else if elf1RangeEnd >= elf2RangeStart && elf1RangeEnd <= elf2RangeEnd {
			count++
		}
	}
	return count
}

func Solution() {
	fmt.Printf("Day 4: Solution 1: %v, Solution 2: %v", solution1(), solution2())
}