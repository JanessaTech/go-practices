package functions

import "fmt"

func sum(a, b int) int {
	return a + b
}

func area(a, b int) int {
	return a * b
}

func cal(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func FunAsMethodParam() {
	r1 := cal(sum, 2, 3)
	fmt.Println("sum(2, 3) = ", r1)
	r2 := cal(area, 2, 3)
	fmt.Println("area(2, 3) = ", r2)
}
