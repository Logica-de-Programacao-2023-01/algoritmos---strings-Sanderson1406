package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Print("Escreva 1 string: ")
	fmt.Scanln(&s1)
	fmt.Print("Escreva a outra string: ")
	fmt.Scanln(&s2)
	if s1 == s2 {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}
}
