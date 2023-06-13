package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string = "Oi bom dia como vai a vida"
	fmt.Println(frase)
	frase = strings.ToUpper(frase)
	fmt.Println(frase)
}
