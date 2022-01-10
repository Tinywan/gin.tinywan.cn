package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j = 10, 20
	fmt.Println("i = ", i, ", j = ", j)

	var (
		name = "Tinywan"
		age  = 24
	)

	fmt.Println("name ", name, ", age = ", age)

	first := "1"
	second := 2
	fmt.Println("first ", first, ", second ", second)

	schoole, lang, color := "QinHua", "PHP", "red"
	fmt.Println("schoole ", schoole, ", lang ", lang, ", color ", color)

	//a, b := 20, 30 // 声明变量a和b
	//fmt.Println("a is", a, "b is", b)
	//
	//b, c := 40, 50 // b已经声明，但c尚未声明
	//fmt.Println("b is", b, "c is", c)
	//
	//b, c = 80, 90 // 给已经声明的变量b和c赋新值
	//fmt.Println("changed b is", b, "c is", c)

	a, b := 20, 30 // 声明a和b
	fmt.Println("a is", a, "b is", b)
	a, bb := 40, 50 // 错误，没有尚未声明的变量
	fmt.Println("a is", a, "bb is", bb)

	aa, cc := 12.90, 543.90
	c := math.Min(aa, cc)
	fmt.Println("c = ", c)
}
