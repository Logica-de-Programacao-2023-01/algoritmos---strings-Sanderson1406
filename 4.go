package main

import "fmt"

func main() {
	var s1 string = ""
	var s2 string = ""
	fmt.Println("Escreva 1 string: ")
	fmt.Scanln(&s1)
	fmt.Println("Escreva a outra string: ")
	fmt.Scanln(&s2)
	if s1 == s2 {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}
}
