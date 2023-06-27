package main

import "fmt"

func main() {
	var s string
	fmt.Print("Infomorme uma string: ")
	fmt.Scan(&s)
	for a := 0; a < len(s)/2; a++ {
		if s[a] != s[len(s)-a-1] {
			fmt.Print("Não é  palidromo")
			break
		} else {
			fmt.Print("É um palidromo")
		}
	}
}
