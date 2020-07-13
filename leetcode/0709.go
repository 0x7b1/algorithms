package main

import (
	"fmt"
	"strings"
)

func toLowerCase(str string) string {
	return strings.ToLower(str)
}

func main() {
	fmt.Println(toLowerCase("LOVElY"))
	fmt.Println(toLowerCase("Hello"))
}
