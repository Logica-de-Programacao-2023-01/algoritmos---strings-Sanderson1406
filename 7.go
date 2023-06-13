package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s1 string
	var tem bool = false
	fmt.Println("Escreva 1 string: ")
	fmt.Scanln(&s1)
	for _, ss := range s1 {
		if unicode.IsDigit(ss) {
			tem = true
			break
		}
	}
	if tem {
		fmt.Println("Há um número")
	} else {
		fmt.Println("Não há um número")
	}
}
