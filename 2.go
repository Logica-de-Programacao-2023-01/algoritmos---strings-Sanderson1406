package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Oi, as vezes a vida é legal"
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
}
