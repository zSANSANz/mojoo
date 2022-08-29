package main

import "fmt"

func selectionSort(arr []float32) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] < arr[j] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}
func main() {
	arr := []float32{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	selectionSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")
}
