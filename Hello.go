package main

import (
	"fmt"
	"home/murray/GoProjects/UnitTestingGoLang/Sorts"
	"home/murray/GoProjects/UnitTestingGoLang/Calculator"
)

func Hello() string{
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
	fmt.Println("Visualizing Bubble Sort")
	fmt.Println(Sorts.BubbleSort([]int{9,1,6,8,7,2,10,3,4,5}))
	fmt.Println("Visualizing Merge Sort")
	fmt.Println(Sorts.MergeSort([]int{9,1,6,8,7,2,10,3,4,5}))
	fmt.Println("Visualizing Quick Sort")
	fmt.Println(Sorts.QuickSort([]int{9,1,6,8,7,2,10,3,4,5}))
	fmt.Println("Calculator Functionality")
	fmt.Println(Calculator.Addition(50,50))
	fmt.Println(Calculator.Multiplication(50,50))
}	