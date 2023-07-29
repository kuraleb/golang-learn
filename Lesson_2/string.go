package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "qwerty"
	b := "text"
	c := "tехt"
	fmt.Println(a[0])
	fmt.Println(" ", len(b), "\n ", len(c))

	aUTF8 := utf8.RuneCountInString(b)
	fmt.Println(" ", aUTF8)
}