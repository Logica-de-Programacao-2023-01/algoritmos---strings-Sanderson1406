package main

import (
	"fmt"
	"strings"
)

func main() {
	var ss string
	fmt.Println("Escreva uma string: ")
	fmt.Scanln(&ss)
	newss := strings.Fields(ss)
	fmt.Println(len(newss))
}
