package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, final string
	fmt.Print("Informe uma string ")
	fmt.Scan(&x)
	final = x
	V := "AEIOUaeiou"
	for _, b := range V {
		final = strings.ReplaceAll(final, string(b), "")
	}
	fmt.Print(final)
}
