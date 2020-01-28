package main

import (
    "fmt"
)

func bubbleSort(items []int) {
    // n is the number of items in our list
    n := len(items)
    // set swapped to true
    sorted := false
    // loop
    for !sorted {
        // set swapped to false
        swapped := false
        // iterate through all of the elements in our list
        for i := 0; i < n-1; i++ {
            // if the current element is greater than the next
            // element, swap them
            if items[i] > items[i+1] {

                // swap values using tuple assignment
                items[i+1], items[i] = items[i], items[i+1]
                // set swapped to true - this is important
                // if the loop ends and swapped is still equal
                // to false, our algorithm will assume the list is
                // fully sorted.
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        // Since the last element is the largest in the slice,
        // we don't need to check it again
        n = n -1
    }
    // finally, print out the sorted list
    fmt.Println(items)
}

func main() {
    toBeSorted := []int{1, 3, 6, 8, 4, 9, 7, 2, 5}
    bubbleSort(toBeSorted)
}