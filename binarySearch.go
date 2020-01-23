// Binary In Golang //
package main 

import "fmt"

func binarySearch(find int, values []int) int {
	low := 0
	high := len(values) - 1


	for low <= high {
		
		mid := (low + high) / 2

		if values[mid] < find {
			low = mid + 1
		} else {
			high = mid - 1
		} 
	}

	if low == len(values) || values[low] != find {
		return -1
	} else {
		return low
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearch(8, values))
} 