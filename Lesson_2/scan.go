package main

import "fmt"

func main() {
	var name string
	var age string
	fmt.Println("Введи имя:")
	fmt.Scan(&name)
	fmt.Println("Введи возраст:")
	fmt.Scan(&age)

	fmt.Println("Привет", name, "\nтебе", age, "лет")
}