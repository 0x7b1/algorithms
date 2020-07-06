package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address,".", "[.]")
}

func main() {
	ip1 := "255.100.50.0"
	fmt.Println(defangIPaddr(ip1))
}
