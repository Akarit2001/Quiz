package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var mem [1000][1000]int

func main() {
	var arr [][]int
	// arr = [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	file, err := os.Open("hard.json")
	notErr(err)
	defer file.Close()
	notErr(json.NewDecoder(file).Decode(&arr))

	fmt.Println(findMaxPath(0, 0, len(arr), len(arr[len(arr)-1]), arr))
}

func findMaxPath(i, j, row, col int, arr [][]int) int {
	if j == col {
		return 0
	}
	if i == row-1 {
		return arr[i][j]
	}

	if mem[i][j] != 0 {
		return mem[i][j]
	}
	mem[i][j] = arr[i][j] + max(findMaxPath(i+1, j, row, col, arr), findMaxPath(i+1, j+1, row, col, arr))
	return mem[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func notErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func greedy(arr [][]int) {
// 	sum := arr[0][0]
// 	maxIndex := 0
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i][maxIndex] > arr[i][maxIndex+1] {
// 			sum += arr[i][maxIndex]
// 		} else {
// 			sum += arr[i][maxIndex+1]
// 			maxIndex++
// 		}

// 	}
// 	fmt.Println(sum)
// }
