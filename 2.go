package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Println("Escreva a string: ")
	fmt.Scanln(&s)
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
}
