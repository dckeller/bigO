package main 

import (
	"fmt"
)

func sort(nums []int, guess int) {
	for _, v := range nums {
		if v == guess {
			fmt.Printf("Found %d \n", guess)
		} else {
			fmt.Println("Could not find")
		}
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sort(array, 1)
}