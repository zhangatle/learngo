package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//定义变量
var (
	aa = 3
	ss = "ttt"
)

// 默认值
func variableZeroValue() {
	var name string
	var age int
	fmt.Printf("%d %q\n", age, name)
}

// 指定初始值
func variableInitalValue() {
	var name string = "zhangatle"
	var age int = 26
	fmt.Println(name, age)
}

// 自动识别类型
func variableTypeDeduction() {
	var name = "zhangatle"
	var age = 26
	fmt.Println(name, age)
}

// 短定义 ：=
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 欧拉
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// 开平方
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 定义常量
func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		php    = iota
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, php)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitalValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
	euler()
	triangle()
	consts()
	enums()
}
