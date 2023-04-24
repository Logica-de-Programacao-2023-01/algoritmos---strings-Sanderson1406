package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Oi, tudo bem, como estao as coisas"
	fmt.Println(s)
	var a, n string
	fmt.Println("Qual caractere você quer trocar?: ")
	fmt.Scanln(&a)
	fmt.Println("Qaul o novo caractere: ")
	fmt.Scanln(&n)
	//for _, ss := range s {}
	s = strings.Replace(s, string(a), string(n), -1)
	fmt.Println(s)
}
