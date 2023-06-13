package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, letra1, letra2 string
	fmt.Println("Escreva uma string")
	fmt.Scanln(&s)
	fmt.Println("Escreva uma letra antiga")
	fmt.Scanln(&letra1)
	fmt.Println("Escreva a nova letra")
	fmt.Scanln(&letra2)
	s = strings.Replace(s, string(letra1), string(letra2), -1)
	fmt.Println(s)
}
