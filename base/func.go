package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数定义
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(9,4)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" , op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d, %d)", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int  {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(2, 4, "x"))
	fmt.Println(div(5,4))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4,
		))
	fmt.Println(sum(2,3,4))
	a, b := 3,4
	a,b = swap(a, b)
	fmt.Println(a,b)
}
