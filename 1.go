package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, result string
	fmt.Println("Escreva uma string ")
	fmt.Scanln(&s)
	for i := 0; i < len(s); i++ {
		result = strings.ToUpper(s)
	}
	fmt.Println(result)
}
