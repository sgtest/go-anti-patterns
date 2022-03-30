package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")
}

func outdatedReplace(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}
