package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	ss := ""
	fmt.Println("Escreva a string: ")
	fmt.Scanln(&s)
	for i := 0; i < len(s); i++ {
		ss = strings.ReplaceAll(s, " ", "")
	}
	fmt.Println(ss)
}
