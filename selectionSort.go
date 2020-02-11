// Selection Sort In Golang //
package main

import "fmt"

func selectionSort(value []int) []int {

	for i := 0; i < len(value); i++ {
		smallest := i
		
		// find the first, second, third, fourth...smallest value
		for j := i; j < len(value); j++ {
			if value[j] < value[smallest] {
				smallest = j
			}
		}

		// Swap using tuple assignment swap smallest with postion of i//
		value[i], value[smallest] = value[smallest], value[i]
	}
	return value
}

func main() {
	arr := []int{5, 2, 1, 8, 5, 0}
	fmt.Println(selectionSort(arr))
}