// Quicksort In Golang //
package main 

import "fmt"

func quicksort(values []int) []int {
	if len(values) < 2 {
		return values				// <--- Base Case
	}

	left := 0
	right := len(values) -1

	// Pick a pivot or starting point //
	pivot := len(values) / 2 // <--- Recursive Case
	fmt.Println(pivot)

	// Move the pivot to the right using tuple assignment //
	values[pivot], values[right] = values[right], values[pivot]

	// Pile elements smaller than the pivot to the left //
  for nums := range values {
  	if values[nums] < values[right] {

  		// Swap using tuple assignment //
    	values[nums], values[left] = values[left], values[nums]
    	left++
  	}
	}

	// Place the pivot after the last smaller element
  values[left], values[right] = values[right], values[left]

  quicksort(values[:left])
  quicksort(values[left+1:])

  return values
}

func main() {
	values := []int{1, 7, 9, 2, 5, 10, 5, 8}
	fmt.Println(quicksort(values))
}