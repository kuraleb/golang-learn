package main

import (
	"fmt"
)

func main() {
	var number1 int
	var number2 int
	fmt.Println("Введите число a")
	fmt.Scan(&number1)
	fmt.Println("Введите число b")
	fmt.Scan(&number2)
	if number1 > number2 {
		fmt.Println("a больше b")
	} else {
		fmt.Println("b больше a")
	}
}