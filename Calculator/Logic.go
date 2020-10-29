package Calculator

import(
	"fmt"
)

func Addition(x int, y int) int{
	return x + y 
}

func Multiplication(x int, y int) int{
	return x * y
}

func Division(x int, y int) int{
	return x / y
}

func Subtraction(x int, y int) int{
	return x - y
}

func main(){
	fmt.Println("Practising Testing on Go - Basic Calculator")
}