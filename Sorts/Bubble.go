package Sorts

import(
	"fmt"
)

func BubbleSort(numbers []int) []int{
	for i := len(numbers); i > 0; i--{
		for j := 1; j < i; j++{
			if numbers[j-1] > numbers[j] {
				temp := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = temp
				fmt.Println(numbers)
			}
		}
	}
	return numbers
}
