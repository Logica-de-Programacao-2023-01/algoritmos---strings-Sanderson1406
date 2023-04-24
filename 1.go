package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Escreva uma string ")
	fmt.Scanln(&s)
	for _, i := range s {
		i++
		result := strings.ToUpper(s)
		fmt.Println(result)
	}

}
