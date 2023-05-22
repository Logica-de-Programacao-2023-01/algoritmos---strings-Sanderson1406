package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = ""
	fmt.Println("Escreva a string: ")
	fmt.Scanln(&s)
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
}
