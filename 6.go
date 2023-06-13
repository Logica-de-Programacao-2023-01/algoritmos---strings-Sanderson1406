package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Oi, como vão as coisas na vida"
	ss := strings.Fields(s)
	quantidade := len(ss)
	fmt.Println("O número de palavras é: ", quantidade)
}
