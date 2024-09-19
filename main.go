package main

import (
	"fmt"
	"main/day1"
	"main/day2"
	"main/day3"
	"main/day4"
	"main/day5"
	"os"
)

func prepScript() {
	for i := 1; i <= 25; i++ {
		folderInfo, err := os.Stat(fmt.Sprintf("day%v", i))
		if os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("day%v", i), os.ModePerm)
			goFile, _ := os.Create(fmt.Sprintf("day%v/day%v.go", i, i))
			mockFile, _ := os.Create(fmt.Sprintf("day%v/mock.txt", i))
			taskFile, _ := os.Create(fmt.Sprintf("day%v/task.txt", i))
			defer goFile.Close()
			defer mockFile.Close()
			defer taskFile.Close()
			fmt.Fprintf(goFile, "package day%v\n\n import (\n\"fmt\"\n)\n\nfunc solution1() int {\n  return 0\n} \n\nfunc solution2() int {\n  return 0\n}\n\nfunc Solution() {\n  fmt.Printf(\"Day %v: Solution 1: %%v, Solution 2: %%v\", solution1(), solution2())\n}", i, i)	
		} else if folderInfo.IsDir() {
			fmt.Printf("Folder day%v already exists\n", i);
		}
	}
}

func main() {
	// prepScript()
	day1.Solution()
	day2.Solution()
	day3.Solution()
	day4.Solution()
	day5.Solution()
}