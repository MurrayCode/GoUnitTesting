package Sorts

import (
	"math/rand"
	"fmt"
)

func QuickSort(numbers []int) []int{
	if len(numbers) < 2 {
        return numbers
    }
	left,right := 0, len(numbers) -1

	pivot := rand.Int() % len(numbers)
	fmt.Println(numbers)
	numbers[pivot], numbers[right] = numbers[right], numbers[pivot]

	for i, _ := range numbers{
		fmt.Println(numbers)
		if numbers[i] < numbers[right]{
			numbers[left], numbers[i] = numbers[i], numbers[left]
			left++
		}
	}
	numbers[left], numbers[right] = numbers[right], numbers[left]
	QuickSort(numbers[:left])
	QuickSort(numbers[left+1:])
	return numbers
}

