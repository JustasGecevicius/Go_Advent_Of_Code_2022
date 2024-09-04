package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type intHeap []int

func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))	
}

func (h *intHeap) Pop() any {
	oldHeap := *h
	heapLength := len(oldHeap)
	removedElement := oldHeap[heapLength-1]
	*h = oldHeap[0 : heapLength - 1]
	return removedElement 
}

func handleAdding(heap *intHeap, count *int) {
			heap.Push(*count)
			for i := len(*heap) - 1; i > 0; i-- {
				if heap.Less(i,i-1) {
					heap.Swap(i,i-1)
				}
			}
			if heap.Len() == 4 {
				heap.Pop()
			}
			*count = 0;
}

// Comment Added
func Solution() {
	file, err := os.Open("day1/task.txt")
	if err != nil { panic(err)}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	calorieHeap := &intHeap{}
	currentCalorieCount := 0

	for scanner.Scan() {
		calorieCount, err := strconv.Atoi(scanner.Text())
		if (err != nil) {
			handleAdding(calorieHeap, &currentCalorieCount)
			continue
		}
		currentCalorieCount += calorieCount
	}

	if currentCalorieCount != 0 {
		handleAdding(calorieHeap, &currentCalorieCount)
	}
	
	if err1 := scanner.Err(); err1 != nil {
		fmt.Print(err1)
	}

	top3Sum := 0
	for _, num := range *calorieHeap {
		top3Sum += num
	}
	fmt.Printf("Day 1: Solution 2: %v", top3Sum)
	fmt.Println()
}
