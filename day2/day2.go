package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getSignScore(x byte) int {
	numx := int(x)
	switch numx {
	case 88:
		return 1
	case 89:
		return 2
	case 90:
		return 3
	default:
		return 0
	}
}

func getTotalScore1(x byte, y byte) int {
	numx := int(x)
	numy := int(y)
	diff := numy - numx;
	switch diff {
	case 23:
		return 3 + getSignScore(y)
	case 24:
		fallthrough
	case 21:
		return 6 + getSignScore(y)
	}
	return 0 + getSignScore(y)
}

func getTotalScore2(x byte, y byte) int {
	RPS := int(x) - 65 // R = 0, P = 1, S = 2
	switch int(y) {
	case 88: // Lose
		return int(math.Mod(float64(RPS + 2), 3)) + 1
	case 89: // Draw
		return 3 + RPS + 1
	case 90: // Win
		return 6 + int(math.Mod(float64(RPS + 1), 3)) + 1
	default:
		return 0
	}
}

func Solution() {	
	file, err := os.Open("day2/task.txt")
	if (err != nil) { panic(err) }
	defer file.Close();
	scanner := bufio.NewScanner(file);
	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		// Solution 1
		total1 += getTotalScore1(line[0], line[2])
		// Solution 2
		total2 += getTotalScore2(line[0], line[2])
	}

	fmt.Printf("Day 2: Solution 1: %v, Solution 2: %v", total1,total2)
	fmt.Println()
}