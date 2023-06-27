package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	numeros := "0123456789"
	fmt.Print("Informe um número: ")
	fmt.Scan(&x)
	tem := false
	for _, J := range numeros {
		if strings.Contains(x, string(J)) {
			tem = true
			break
		}
	}
	if tem {
		fmt.Print("tem número")
	} else {
		fmt.Print("Não tem número")
	}
}
